package api

import (
	"encoding/json"
	"net/http"

	"github.com/mirno/taskmanagement/internal/usecase"
)

type TaskHandler struct {
	service *usecase.TaskService
}

func NewTaskHandler(service *usecase.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := h.service.CreateTask(input.Title, input.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// TODO Add functions for: Update, Retrieve, Delete of Tasks
