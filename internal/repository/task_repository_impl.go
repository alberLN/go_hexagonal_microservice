package repository

import (
	"context"
	"example/microservice/internal/domain"

	"github.com/jinzhu/gorm"
)

type TaskRepositoryImpl struct {
	db *gorm.DB
}

func NewTaskRepositoryImpl(db *gorm.DB) *TaskRepositoryImpl {
	return &TaskRepositoryImpl{db: db}
}

func (t *TaskRepositoryImpl) GetById(ctx context.Context, id int) (*domain.Task, error) {
	task := &domain.Task{}
	if err := t.db.First(task, id).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (t *TaskRepositoryImpl) GetAll(ctx context.Context) ([]*domain.Task, error) {
	var tasks []*domain.Task
	t.db.Find(&tasks)
	err := t.db.Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (t *TaskRepositoryImpl) Create(ctx context.Context, task *domain.Task) (*domain.Task, error) {
	if err := t.db.Create(&task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (t *TaskRepositoryImpl) Update(ctx context.Context, task *domain.Task) error {
	if err := t.db.Save(task).Error; err != nil {
		return err
	}
	return nil
}

func (t *TaskRepositoryImpl) Delete(ctx context.Context, id int) error {
	if err := t.db.Delete(&domain.Task{}, id).Error; err != nil {
		return err
	}
	return nil
}
