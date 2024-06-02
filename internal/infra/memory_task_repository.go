package infra

import (
	"errors"

	"github.com/mirno/taskmanagement/internal/domain"
)

type MemoryTaskRepository struct {
	tasks map[string]*domain.Task
}

func NewMemoryTaskRepository() *MemoryTaskRepository {
	return &MemoryTaskRepository{
		tasks: make(map[string]*domain.Task),
	}
}

func (r *MemoryTaskRepository) Save(task *domain.Task) error {
	r.tasks[task.ID.String()] = task
	return nil
}

func (r *MemoryTaskRepository) GetByID(id string) (*domain.Task, error) {
	task, exists := r.tasks[id]
	if !exists {
		return nil, errors.New("task not found")
	}

	return task, nil
}

func (r *MemoryTaskRepository) GetAll() ([]*domain.Task, error) { return nil, domain.ErrUnimplemented }
func (r *MemoryTaskRepository) Update(task *domain.Task) error  { return domain.ErrUnimplemented }
func (r *MemoryTaskRepository) Delete(id string) error          { return domain.ErrUnimplemented }
