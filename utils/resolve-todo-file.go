package utils

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/masputrawae/go-todo/pkg/model"
)

func LoadTodo(fp string) ([]model.Todo, error) {
	var data []model.Todo
	file, err := resolveFile(fp)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		if err.Error() == "EOF" {
			return data, json.NewEncoder(file).Encode(&data)
		}
		return nil, err
	}
	return data, nil
}

func SaveTodo(fp string, data []model.Todo) error {
	file, err := resolveFile(fp)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(data)
}

func resolveFile(fp string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(fp), 0755); err != nil {
		return nil, err
	}
	file, err := os.OpenFile(fp, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return file, nil
}
