package db

import "context"

type DBInterface interface {
	AddTask(ctx context.Context, dto TasksDTO) (Task, error)
	Tasks(ctx context.Context) ([]TasksResp, error)
	UpdateTask(ctx context.Context, dto TasksDTO, id string) (TasksResp, error)
	DeleteTask(ctx context.Context, id string) (bool, error)
	CompleteTask(ctx context.Context, dto CompleteDTO, id string) (TasksResp, error)
}

type CompleteDTO struct {
	Completed bool `json:"completed"`
}

type TasksDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
}

type TasksResp struct {
	Title       string
	Description string
	DueDate     string
	Overdue     bool
	Completed   bool
}

type Task struct {
	Id          int
	Title       string
	Description string
	DueDate     string
	Overdue     bool
	Completed   bool
}
