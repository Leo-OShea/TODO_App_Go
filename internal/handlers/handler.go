package handlers

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"

	"leo.com/m/internal/models"
)

var tasks = []models.Task{
	{ID: 1, Title: "Task 1", Completed: false},
	{ID: 2, Title: "Task 2", Completed: false},
}

// curl localhost:8080
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

// (i)
// curl localhost:8080/tasks
func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "To-Do List:")
	fmt.Fprintln(w, "--------------")

	for i := 0; i < len(tasks); i++ {
		tmp := i + 1
		if tasks[i].Completed {
			fmt.Fprintf(w, "[x] %d. %s\n", tmp, tasks[i].Title)
		} else {
			fmt.Fprintf(w, "[ ] %d. %s\n", tmp, tasks[i].Title)
		}
	}
}

// (ii)
// curl localhost:8080/tasks/{id}
func GetSpecificTaskHandler(w http.ResponseWriter, r *http.Request) {
	
	id := r.PathValue("id")

	i, err := strconv.Atoi(id)
	i -= 1
	if err != nil || i < 0 || i >= len(tasks) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid task ID: %s", id)
	} else {
		fmt.Fprintf(w, "Task with ID %s:\n", id)
		fmt.Fprintf(w, "%s", tasks[i].Title)
	}
}

// (iii)
// curl -X POST localhost:8080/addTask/{title}
func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	title := r.PathValue("title")
	if title == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Task title cannot be empty.")
		return
	} else {
		tasks = append(tasks, models.Task{
		ID:        len(tasks) + 1,
		Title:     title,
		Completed: false,
	})
	fmt.Fprintf(w, "Added task: %s\n", title)
	}
}

// (iv)
// curl -X PUT localhost:8080/tasks/{id}
func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	i, err := strconv.Atoi(id)
	i -= 1
	if err != nil || i < 0 || i >= len(tasks) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid task ID: %s", id)
	} else {
		tasks[i].Completed = !tasks[i].Completed
		fmt.Fprintf(w, "Updated task with ID: %s", id)
	}
}

// (v)
// curl -X DELETE localhost:8080/tasks/{id}
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	i, err := strconv.Atoi(id)
	i -= 1
	if len(tasks) == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No tasks to delete")
	} else if err != nil || i < 0 || i >= len(tasks) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid task ID: %s", id)
	} else {
		// deletes a range of elements from the slice (which is why it needs two arguments)
		tasks = slices.Delete(tasks, i, i+1)
		fmt.Fprintf(w, "Deleted task with ID: %s", id)
	}
}
