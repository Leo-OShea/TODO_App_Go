package main

import (
	"net/http"
	"net/http/httptest"
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
	if rr.Body.String() != expected {
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

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.GetTasksHandler(w, r)
	})

	handler.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status OK; got %v", rr.Code)
	}

	expected := "To-Do List:\n--------------\nNo tasks available."
	if rr.Body.String() != expected {
		t.Errorf("\nUnexpected response body.\nGot: \n%v\n\nExpected: \n%v", rr.Body.String(), expected)
	}
}
