package main

import (
	"fmt"
	"net/http"

	"leo.com/m/internal/handlers"
	// "leo.com/m/internal/models"
)

func main() {
	fmt.Println("Todo List")

	mux := http.NewServeMux()

	// Say hello!!
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HelloHandler(w, r)
	})

	// GET /tasks
	// Return all tasks
	mux.HandleFunc("GET /tasks", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTasksHandler(w, r)
	})

	// GET /tasks/{id}
	// Return specific task
	mux.HandleFunc("GET /tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetSpecificTaskHandler(w, r)
	})

	// POST /tasks
	// Add a task
	mux.HandleFunc("POST /addTask/{title}", func(w http.ResponseWriter, r *http.Request) {
		handlers.AddTaskHandler(w, r)
	})

	// PUT /tasks/{id}
	// Update a task as completed/uncompleted
	mux.HandleFunc("PUT /tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateTaskHandler(w, r)
	})

	// DELETE /tasks/{id}
	// Delete a task
	mux.HandleFunc("DELETE /tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteTaskHandler(w, r)
	})

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
