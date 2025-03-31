package main

import (
	"fmt"
	"net/http"

	"leo.com/m/internal/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HelloHandler)

	fmt.Println("Server is listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
