package storer

// responsible for storing and retrieving tasks to/from todos.json

import (
	"encoding/json"
	"fmt"
	"os"

	"leo.com/m/internal/models"
)

const todosFile = "./internal/data/todos.json"

// saves tasks to todos.json
// converts slice to JSON
func SaveTodos(tasks []models.Task) {

	b, err := json.Marshal(tasks)
	if err != nil {
		fmt.Printf("Failed to marshal tasks: %v\n", err)
		return
	}

	// write to file
	err = os.WriteFile(todosFile, b, 0644)
	if err != nil {
		fmt.Printf("Failed to write tasks to file: %v\n", err)
		return
	}

}

// loads tasks from todos.json
// returns a slice of tasks
func LoadTodos() []models.Task {

	// read the file
	file, err := os.ReadFile(todosFile)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		return []models.Task{}
	}

	var tasks []models.Task
	if len(file) > 0 {
		// second arg is where the data is stored
		err = json.Unmarshal(file, &tasks)
		if err != nil {
			fmt.Printf("Failed to unmarshal JSON: %v\n", err)
			return []models.Task{}
		}
	} else {
		fmt.Printf("File is empty")
		return []models.Task{}
	}

	fmt.Printf("Loaded tasks")

	return tasks
}
