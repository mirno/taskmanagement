package main

import (
	"log"
	"net/http"
	"time"

	"github.com/mirno/taskmanagement/internal/infra"
	"github.com/mirno/taskmanagement/internal/interfaces/api"
	"github.com/mirno/taskmanagement/internal/usecase"
)

func main() {
	repo := infra.NewMemoryTaskRepository()
	service := usecase.NewTaskService(repo) // insert the implementation of infrastructure / driver into the Service which needs to met the interface
	handler := api.NewTaskHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", handler.CreateTask)
	// Additional routes...

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,  //  Maximum duration for reading the entire request, including the body.
		WriteTimeout: 10 * time.Second, // Maximum duration before timing out writes of the response.
		IdleTimeout:  15 * time.Second, // Maximum amount of time to wait for the next request when keep-alives are enabled.
	}

	log.Fatal(server.ListenAndServe())
}
