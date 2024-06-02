package utils

import (
	"github.com/google/uuid"
)

func GenerateID() uuid.UUID {
	id := uuid.New()
	return id
}
