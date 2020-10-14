package main

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	pbActivity "github.com/homma509/9task/proto/activity"
	pbProject "github.com/homma509/9task/proto/project"
	pbTask "github.com/homma509/9task/proto/task"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TaskService Taskサービス構造体
type TaskService struct {
	store          Store
	activityClient pbActivity.ActivityServiceClient
	projectClient  pbProject.ProjectServiceClient
}

// CreateTask Taskの登録
func (s *TaskService) CreateTask(
	ctx context.Context,
	req *pbTask.CreateTaskRequest,
) (*pbTask.CreateTaskResponse, error) {
	if req.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument,
			"empty task name")
	}

	resp, err := s.projectClient.FindProject(ctx,
		&pbProject.FindProjectRequest{ProjectId: req.GetProjectId()})
	if err != nil {
		return nil, status.Error(
			codes.NotFound, err.Error())
	}
	// TODO Placeholder ID
	userID := uint64(1)
	now := ptypes.TimestampNow()
	task, err := s.store.CreateTask(&pbTask.Task{
		Name:      req.GetName(),
		Status:    pbTask.Status_WAITING,
		UserId:    userID,
		ProjectId: resp.Project.GetId(),
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return nil, status.Error(
			codes.InvalidArgument, err.Error(),
		)
	}

	content := &pbActivity.CreateTaskContent{
		TaskId:   task.GetId(),
		TaskName: task.GetName(),
	}
	any, err := ptypes.MarshalAny(content)
	if err != nil {
		return nil, status.Error(
			codes.InvalidArgument, err.Error(),
		)
	}
	if _, err := s.activityClient.CreateActivity(ctx,
		&pbActivity.CreateActivityRequest{
			Content: any,
		}); err != nil {
		return nil, err
	}
	return &pbTask.CreateTaskResponse{Task: task}, nil
}

// FindTasks ユーザーに関連するタスクの検索
func (s *TaskService) FindTasks(
	ctx context.Context,
	_ *empty.Empty,
) (*pbTask.FindTasksResponse, error) {
	// TODO Placeholder ID
	userID := uint64(1)
	tasks, err := s.store.FindTasks(userID)
	if err != nil {
		return nil, status.Error(
			codes.InvalidArgument, err.Error())
	}
	return &pbTask.FindTasksResponse{Tasks: tasks}, nil
}