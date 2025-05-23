package handlers

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"

	"leo.com/m/internal/models"
	"leo.com/m/internal/storer"
)

const todosFile = "./internal/data/todos.json"

var tasks = storer.LoadTodos(todosFile)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Empty fields are not allowed.")
	fmt.Fprintln(w, "Please refer to the provided endpoints.")
}

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {

	tasks = storer.LoadTodos(todosFile)

	if len(tasks) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "No tasks available.")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "To-Do List:\n--------------")

		for i := 0; i < len(tasks); i++ {

			tmp := i + 1
			if tasks[i].Completed {
				fmt.Fprintf(w, "[x] %d. %s\n", tmp, tasks[i].Title)
			} else {
				fmt.Fprintf(w, "[ ] %d. %s\n", tmp, tasks[i].Title)
			}
		}
	}
}

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

	storer.SaveTodos(todosFile, tasks)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	i, err := strconv.Atoi(id)
	i -= 1
	if err != nil || i < 0 || i >= len(tasks) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid task ID: %s", id)
	} else {
		w.WriteHeader(http.StatusOK)
		tasks[i].Completed = !tasks[i].Completed
		fmt.Fprintf(w, "Updated task with ID: %s", id)
	}

	storer.SaveTodos(todosFile, tasks)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	tasks = storer.LoadTodos(todosFile)

	i, err := strconv.Atoi(id)
	i -= 1
	if len(tasks) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "No tasks to delete")
	} else if err != nil || i < 0 || i >= len(tasks) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid task ID: %s", id)
	} else {
		tasks = slices.Delete(tasks, i, i+1)
		fmt.Fprintf(w, "Deleted task with ID: %s", id)
	}

	storer.SaveTodos(todosFile, tasks)
}
