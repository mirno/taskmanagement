package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrUnimplemented = errors.New("feature not implemented")
	ErrNotSupported  = errors.New("operation not supported")
)

type Task struct {
	ID          uuid.UUID
	Title       string
	Description string
	Status      string
}

type TaskRepository interface {
	Save(task *Task) error
	GetByID(id string) (*Task, error)
	GetAll() ([]*Task, error)
	Update(task *Task) error
	Delete(id string) error
}
