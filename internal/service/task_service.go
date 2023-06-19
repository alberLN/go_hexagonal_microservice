package service

import (
	"context"
	"example/microservice/internal/domain"
)

type TaskService interface {
	CreateTask(ctx context.Context, taskRQ *domain.TaskRequest) (*domain.Task, error)
	GetTaskById(ctx context.Context, taskId int) (*domain.Task, error)
	UpdateTask(ctx context.Context, taskId int, taskRQ *domain.TaskRequest) error
	DeleteTask(ctx context.Context, taskId int) error
	GetAll(ctx context.Context) ([]*domain.Task, error)
}
