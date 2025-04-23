package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"leo.com/m/internal/handlers"
	"leo.com/m/internal/storer"
)

const todosFile = "../internal/data/todos.json"

var tasks = storer.LoadTodos(todosFile)

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
		t.Errorf("Got unexpected status %v", rr.Code)
	}

	expected := "Empty fields are not allowed.\nPlease refer to the provided endpoints.\n"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("Unexpected response body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestGetTasks(t *testing.T) {

	req, err := http.NewRequest("GET", "/todos", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTasksHandler(w, r)
	})

	handler.ServeHTTP(rr, req)

	if len(tasks) == 0 {
		if rr.Code != http.StatusBadRequest {
			t.Errorf("Got unexpected status %v", rr.Code)
		}

		expected := "No tasks available."
		if !strings.Contains(rr.Body.String(), expected) {
			t.Errorf("\nUnexpected response body.\nGot: \n%v\n\nExpected: \n%v", rr.Body.String(), expected)
		}
	} else { // Non-empty todo list
		if rr.Code != http.StatusOK {
			t.Errorf("Got unexpected status %v", rr.Code)
		}

		expected := "To-Do List:\n--------------"
		if !strings.Contains(rr.Body.String(), expected) {
			t.Errorf("\nUnexpected response body.\nGot: \n%v\n\nExpected: \n%v", rr.Body.String(), expected)
		}
	}
}

func TestGetSpecificTask(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /todos/{id}", handlers.GetSpecificTaskHandler)

	req, err := http.NewRequest("GET", "/todos/1", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	if len(tasks) == 0 {
		if rr.Code != http.StatusBadRequest {
			t.Errorf("Got unexpected status %v", rr.Code)
		}

		expected := "Invalid task ID: 1"
		if !strings.Contains(rr.Body.String(), expected) {
			t.Errorf("\nUnexpected response body.\nGot: \n%v\n\nExpected: \n%v", rr.Body.String(), expected)
		}
	} else {
		if rr.Code != http.StatusOK {
			t.Errorf("Got unexpected status %v", rr.Code)
		}

		expected := "Task with ID 1:\nTestTitle"
		if !strings.Contains(rr.Body.String(), expected) {
			t.Errorf("\nUnexpected response body.\nGot: \n%v\n\nExpected: \n%v", rr.Body.String(), expected)
		}
	}
}

func TestUpdateTask(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("PUT /todos/{id}", handlers.UpdateTaskHandler)

	req, err := http.NewRequest("PUT", "/todos/1", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	if len(tasks) == 0 {
		if rr.Code != http.StatusBadRequest {
			t.Errorf("Got unexpected status %v", rr.Code)
		}

		expected := "Invalid task ID: 1"
		if !strings.Contains(rr.Body.String(), expected) {
			t.Errorf("\nUnexpected response body.\nGot: \n%v\n\nExpected: \n%v", rr.Body.String(), expected)
		}
	} else {
		if rr.Code != http.StatusOK {
			t.Errorf("Got unexpected status %v", rr.Code)
		}

		expected := "Updated task with ID: 1"
		if !strings.Contains(rr.Body.String(), expected) {
			t.Errorf("\nUnexpected response body.\nGot: \n%v\n\nExpected: \n%v", rr.Body.String(), expected)
		}
	}
}

func TestDeleteTask(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("DELETE /todos/{id}", handlers.DeleteTaskHandler)

	req, err := http.NewRequest("DELETE", "/todos/1", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	if len(tasks) == 0 {
		if rr.Code != http.StatusBadRequest {
			t.Errorf("Got unexpected status %v", rr.Code)
		}

		expected := "No tasks to delete"
		if !strings.Contains(rr.Body.String(), expected) {
			t.Errorf("\nUnexpected response body.\nGot: \n%v\n\nExpected: \n%v", rr.Body.String(), expected)
		}
	} else {
		if rr.Code != http.StatusOK {
			t.Errorf("Got unexpected status %v", rr.Code)
		}

		expected := "Deleted task with ID: 1"
		if !strings.Contains(rr.Body.String(), expected) {
			t.Errorf("\nUnexpected response body.\nGot: \n%v\n\nExpected: \n%v", rr.Body.String(), expected)
		}
	}
}

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
		t.Errorf("Got unexpected status %v", rr.Code)
	}

	expected := "Added task: TestTitle\n"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("\nUnexpected response body.\nGot: \n%v\n\nExpected: \n%v", rr.Body.String(), expected)
	}
}
