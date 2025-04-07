package main

import (
	"fmt"
	"net/http"

	"leo.com/m/internal/handlers"
	// "leo.com/m/internal/models"
)

func main() {
	fmt.Println("--------------------------------")
	fmt.Println("Todo List")
	fmt.Println("--------------------------------")
	fmt.Println("Server started at localhost:8080")
	fmt.Println("--------------------------------")
	fmt.Println("Available endpoints:")
	fmt.Println("- GET /todos")
	fmt.Println("- GET /todos/{id}")
	fmt.Println("- POST /todos/{title}")
	fmt.Println("- PUT /todos/{id}")
	fmt.Println("- DELETE /todos/{id}")
	fmt.Println("--------------------------------")

	mux := http.NewServeMux()

	// Say hello!!
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HelloHandler(w, r)
	})

	// GET /todos
	// Return all tasks
	mux.HandleFunc("GET /todos", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTasksHandler(w, r)
	})

	// GET /todos/{id}
	// Return specific task
	mux.HandleFunc("GET /todos/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetSpecificTaskHandler(w, r)
	})

	// POST /todos/{title}
	// Add a task
	mux.HandleFunc("POST /todos/{title}", func(w http.ResponseWriter, r *http.Request) {
		handlers.AddTaskHandler(w, r)
	})

	// PUT /todos/{id}
	// Update a task as completed/uncompleted
	mux.HandleFunc("PUT /todos/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateTaskHandler(w, r)
	})

	// DELETE /todos/{id}
	// Delete a task
	mux.HandleFunc("DELETE /todos/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteTaskHandler(w, r)
	})

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
