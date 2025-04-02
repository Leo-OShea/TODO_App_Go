package main

import (
	"fmt"
	"net/http"

	"leo.com/m/internal/handlers"
	"leo.com/m/internal/models"
)

var tasks = []models.Task{
	{ID: 1, Title: "Task 1", Completed: false},
	{ID: 2, Title: "Task 2", Completed: false},
}

func main() {
	fmt.Println("Todo List")

	mux := http.NewServeMux()

	// Say hello!!
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HelloHandler(w, r)
	})

	// Return all tasks
	mux.HandleFunc("GET /tasks", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTasksHandler(w, r)
	})

	// Return specific task
	mux.HandleFunc("GET /tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetSpecificTaskHandler(w, r)
	})

	// Add a task
	mux.HandleFunc("POST /addTask/{title}", func(w http.ResponseWriter, r *http.Request) {
		// handlers.AddTaskHandler(w, r)
		title := r.PathValue("title")
		newTask := models.Task{ID: len(tasks) + 1, Title: title, Completed: false}
		tasks = append(tasks, newTask)
		fmt.Fprintf(w, "Added task: %s", title)
	})

	// Update a task as completed/uncompleted
	mux.HandleFunc("PUT /tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		for i, task := range tasks {
			if fmt.Sprintf("%d", task.ID) == id {
				tasks[i].Completed = true
				fmt.Fprintf(w, "Updated task with ID: %s", id)
				return
			}
		}
		http.NotFound(w, r)
	})

	// Delete a task
	mux.HandleFunc("DELETE /tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "Deleted task with ID: %s", id)
	})

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
