package main

import (
	pbTask "github.com/homma509/9task/proto/task"
	"github.com/homma509/9task/shared/inmemory"
)

// Store Task Repository Interface
type Store interface {
	CreateTask(task *pbTask.Task) (*pbTask.Task, error)
	FindTask(taskID, userID uint64) (*pbTask.Task, error)
	FindTasks(userID uint64) ([]*pbTask.Task, error)
	FindProjectTasks(projectID, userID uint64) ([]*pbTask.Task, error)
	updateTask(task *pbTask.Task) (*pbTask.Task, error)
}

// StoreOnMemory タスクのMap構造体
type StoreOnMemory struct {
	tasks *inmemory.IndexMap
}

// NewStoreOnMemory タスクのMap構造体の生成
func NewStoreOnMemory() *StoreOnMemory {
	return &StoreOnMemory{inmemory.NewIndexMap()}
}
