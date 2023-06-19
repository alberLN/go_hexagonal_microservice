package service

import (
	"context"
	"example/microservice/internal/domain"
	"example/microservice/internal/repository"
)

type taskService struct {
	taskRepository repository.TaskRepository
}

func NewTaskServiceImpl(taskRespo repository.TaskRepository) TaskService {
	return &taskService{taskRepository: taskRespo}
}

func (t *taskService) CreateTask(ctx context.Context, taskRQ *domain.TaskRequest) (*domain.Task, error) {

	task := &domain.Task{
		Title:       taskRQ.Title,
		Description: taskRQ.Description,
		Priority:    taskRQ.Priority,
	}

	task, err := t.taskRepository.Create(ctx, task)

	return task, err
}

func (t *taskService) GetTaskById(ctx context.Context, id int) (*domain.Task, error) {
	task, err := t.taskRepository.GetById(ctx, id)
	return task, err
}

func (t *taskService) UpdateTask(ctx context.Context, taskId int, taskRQ *domain.TaskRequest) error {
	task, err := t.taskRepository.GetById(ctx, taskId)
	if err != nil {
		return err
	}

	task.Title = taskRQ.Title
	task.Description = taskRQ.Description
	task.Priority = taskRQ.Priority

	err = t.taskRepository.Update(ctx, task)
	return err
}

func (t *taskService) DeleteTask(ctx context.Context, taskId int) error {
	_, err := t.taskRepository.GetById(ctx, taskId)
	if err != nil {
		return err
	}
	err = t.taskRepository.Delete(ctx, taskId)
	return err
}

func (t *taskService) GetAll(ctx context.Context) ([]*domain.Task, error) {
	task, err := t.taskRepository.GetAll(ctx)
	return task, err
}
