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
	fmt.Println("- curl localhost:8080/todos")
	fmt.Println("- curl localhost:8080/todos/{id}")
	fmt.Println("- curl -X POST localhost:8080/todos/{title}")
	fmt.Println("- curl -X PUT localhost:8080/todos/{id}")
	fmt.Println("- curl -X DELETE localhost:8080/todos/{id}")
	fmt.Println("--------------------------------")

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HelloHandler(w, r)
	})

	mux.HandleFunc("GET /todos", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTasksHandler(w, r)
	})

	mux.HandleFunc("GET /todos/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetSpecificTaskHandler(w, r)
	})

	mux.HandleFunc("POST /todos/{title}", func(w http.ResponseWriter, r *http.Request) {
		handlers.AddTaskHandler(w, r)
	})

	mux.HandleFunc("PUT /todos/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateTaskHandler(w, r)
	})

	mux.HandleFunc("DELETE /todos/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteTaskHandler(w, r)
	})

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
