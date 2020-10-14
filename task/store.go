package main

import (
	pbTask "github.com/homma509/9task/proto/task"
)

// Store Task Repository Interface
type Store interface {
	CreateTask(task *pbTask.Task) (*pbTask.Task, error)
	FindTask(taskID, userID uint64) (*pbTask.Task, error)
	FindTasks(userID uint64) ([]*pbTask.Task, error)
	FindProjectTasks(projectID, userID uint64) ([]*pbTask.Task, error)
	updateTask(task *pbTask.Task) (*pbTask.Task, error)
}
