package wtf

import (
	"context"
	"time"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TaskService interface {
	CreateTask(ctx context.Context, task *Task) error
	UpdateTask(ctx context.Context, id int, upd UpdateTask) (*Task, error)
}

type UpdateTask struct {
	Title       *string
	Description *string
	Completed   *bool
}
