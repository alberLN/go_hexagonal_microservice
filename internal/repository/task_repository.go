package repository

import (
	"context"
	"example/microservice/internal/domain"
)

type TaskRepository interface {
	GetById(ctx context.Context, id int) (*domain.Task, error)
	GetAll(ctx context.Context) ([]*domain.Task, error)
	Create(ctx context.Context, order *domain.Task) (*domain.Task, error)
	Update(ctx context.Context, order *domain.Task) error
	Delete(ctx context.Context, id int) error
}
