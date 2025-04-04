package handlers

import (
	"fmt"
	"net/http"

	"leo.com/m/internal/models"
)

var tasks = []models.Task{
	{ID: 1, Title: "Task 1", Completed: false},
	{ID: 2, Title: "Task 2", Completed: false},
}

// curl localhost:8080
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello, World!")
}

// (i)
// curl localhost:8080/tasks
func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get Tasks")
}

// (ii)
// curl localhost:8080/tasks/{id}
func GetSpecificTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "Return task with ID: %s", id)
}

// (iii)
// curl -X POST localhost:8080/addTask/{title}
func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	title := r.PathValue("title")
	newTask := models.Task{ID: len(tasks) + 1, Title: title, Completed: false}
	tasks = append(tasks, newTask)
	fmt.Fprintf(w, "Added task: %s", title)
}

// (iv)
// curl -X PUT localhost:8080/tasks/{id}
func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "Updated task with ID: %s", id)

}

// (v)
// curl -X DELETE localhost:8080/tasks/{id}
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "Deleted task with ID: %s", id)
}
