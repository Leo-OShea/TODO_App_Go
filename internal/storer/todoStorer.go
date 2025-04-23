package storer

import (
	"encoding/json"
	"fmt"
	"os"

	"leo.com/m/internal/models"
)

func SaveTodos(todosFile string, tasks []models.Task) {

	b, err := json.Marshal(tasks)
	if err != nil {
		fmt.Printf("Failed to marshal tasks: %v\n", err)
		return
	}

	err = os.WriteFile(todosFile, b, 0644)
	if err != nil {
		fmt.Printf("Failed to write tasks to file: %v\n", err)
		return
	}

}

func LoadTodos(todosFile string) []models.Task {

	file, err := os.ReadFile(todosFile)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		return []models.Task{}
	}

	var tasks []models.Task
	if len(file) > 0 {
		err = json.Unmarshal(file, &tasks)
		if err != nil {
			fmt.Printf("Failed to unmarshal JSON: %v\n", err)
			return []models.Task{}
		}
	} else {
		fmt.Printf("File is empty")
		return []models.Task{}
	}

	return tasks
}
