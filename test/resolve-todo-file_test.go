package test

import (
	"github.com/masputrawae/go-todo/pkg/model"
	"github.com/masputrawae/go-todo/utils"
	"os"
	"testing"
)

func removeFile(fp string) {
	os.Remove(fp)
}

func TestLoad(t *testing.T) {
	filePath := "example.json"
	_, err := utils.LoadTodo(filePath)
	if err != nil {
		t.Error(err.Error())
		return
	}
	removeFile(filePath)
}

func TestSave(t *testing.T) {
	filePath := "example.json"
	data, err := utils.LoadTodo(filePath)
	if err != nil {
		t.Error(err.Error())
		return
	}
	data = append(data, model.Todo{ID: 1, Task: "Task 1"})
	data = append(data, model.Todo{ID: 2, Task: "Task 2"})
	if err := utils.SaveTodo(filePath, data); err != nil {
		t.Error(err.Error())
		return
	}
	removeFile(filePath)
}
