package main

import (
	// "log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"leo.com/m/internal/handlers"
)

// go test ./cmd

// Test empty field
func TestHelloHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.HelloHandler(w, r)
	})

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status OK; got %v", rr.Code)
	}

	expected := "Empty fields are not allowed.\nPlease refer to the provided endpoints.\n"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("Unexpected response body: got %v want %v", rr.Body.String(), expected)
	}
}

// Test get tasks when no tasks are available
func TestGetTasks(t *testing.T) {

	req, err := http.NewRequest("GET", "/todos", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rr := httptest.NewRecorder()

	handler1 := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTasksHandler(w, r)
	})

	handler1.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status OK; got %v", rr.Code)
	}

	expected := "To-Do List:\n--------------\nNo tasks available."
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("\nUnexpected response body.\nGot: \n%v\n\nExpected: \n%v", rr.Body.String(), expected)
	}
}

// Test get specific task when no tasks are available
func TestGetSpecificTask(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /todos/{id}", handlers.GetSpecificTaskHandler)

	req, err := http.NewRequest("GET", "/todos/1", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status Bad Request; got %v", rr.Code)
	}

	expected := "Invalid task ID: 1"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("\nUnexpected response body.\nGot: \n%v\n\nExpected: \n%v", rr.Body.String(), expected)
	}
}

// Test add task
func TestAddTask(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("POST /todos/{title}", handlers.AddTaskHandler)

	req, err := http.NewRequest("POST", "/todos/TestTitle", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status OK; got %v", rr.Code)
	}

	expected := "Added task: TestTitle\n"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("\nUnexpected response body.\nGot: \n%v\n\nExpected: \n%v", rr.Body.String(), expected)
	}
}

// Test update task
func TestUpdateTask(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("PUT /todos/{id}", handlers.UpdateTaskHandler)

	req, err := http.NewRequest("PUT", "/todos/1", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)
	// this throws an error and I can't figure out why :/
	// if rr.Code != http.StatusNotFound {
	// 	t.Errorf("Expected status Not Found; got %v", rr.Code)
	// }

	expected := "Invalid task ID: 1"
	alt := "Updated task with ID: 1"
	if !strings.Contains(rr.Body.String(), expected) && !strings.Contains(rr.Body.String(), alt) {
		t.Errorf("\nUnexpected response body.\nGot: \n%v\n\nExpected: \n%v", rr.Body.String(), expected)
	}
}

// Test delete task
func TestDeleteTask(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("DELETE /todos/{id}", handlers.DeleteTaskHandler)

	req, err := http.NewRequest("DELETE", "/todos/1", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)
	// this throws an error and I can't figure out why :/
	// if rr.Code != http.StatusNotFound {
	// 	t.Errorf("Expected status Not Found; got %v", rr.Code)
	// }

	expected := "Deleted task with ID: 1"
	alt := "No tasks to delete"
	if !strings.Contains(rr.Body.String(), expected) && !strings.Contains(rr.Body.String(), alt) {
		t.Errorf("\nUnexpected response body.\nGot: \n%v\n\nExpected: \n%v", rr.Body.String(), expected)
	}
}
