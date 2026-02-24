package utils

import (
	"encoding/json"
	"os"

	"github.com/masputrawae/todo-cli/pkg/model"
)

func LoadTodo(fileName string) ([]model.Todo, error) {
	var data []model.Todo
	f, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return data, json.Unmarshal(f, &data)
}

// save data
func SaveData(fileName string, data []model.Todo) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, jsonData, 0644)
}
