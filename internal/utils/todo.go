package utils

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/masputrawae/todo-cli/internal/model"
)

func LoadTodo(fp string) ([]model.Todo, error) {
	var data []model.Todo
	file, err := ResolveFile(fp)
	if err != nil {
		return data, err
	}
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		if err.Error() != "EOF" {
			return data, err
		}
		if err = json.NewEncoder(file).Encode(&data); err != nil {
			return data, err
		}
	}
	return data, file.Close()
}

// save todo
func SaveTodo(fp string, data []model.Todo) error {
	file, err := ResolveFile(fp)
	if err != nil {
		log.Fatal(err)
	}
	if err = json.NewEncoder(file).Encode(&data); err != nil {
		return err
	}
	return file.Close()
}

// normalize text
func NormalizeText(s string) string {
	return strings.ReplaceAll(strings.ToLower(s), " ", "-")
}

// same text
func TextIsSame(s, i string) bool {
	return NormalizeText(s) == NormalizeText(i)
}
