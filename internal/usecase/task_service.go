package usecase

import (
	"fmt"

	"github.com/mirno/taskmanagement/internal/domain"
	"github.com/mirno/taskmanagement/pkg/utils"
)

type TaskService struct {
	repo domain.TaskRepository
}

func NewTaskService(repo domain.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(title, description string) (*domain.Task, error) {
	task := &domain.Task{
		ID:          utils.GenerateID(),
		Title:       title,
		Description: description,
		Status:      "pending",
	}

	err := s.repo.Save(task)
	if err != nil {
		return nil, fmt.Errorf("unable to save task: %w", err)
	}

	return task, nil
}

// TODO Add functions for: Update, Retrieve, Delete of Tasks
