package handlers

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello, World!")
}

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get Tasks")

}

func GetSpecificTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "Return task with ID: %s", id)
}

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Add Task"))
}
