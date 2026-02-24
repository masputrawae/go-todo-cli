package main

import (
	"log"

	"github.com/masputrawae/todo-cli/pkg/model"
	"github.com/masputrawae/todo-cli/pkg/repo"
	"github.com/masputrawae/todo-cli/pkg/utils"
)

func main() {
	file, err := utils.LoadTodo("dummy.json")
	if err != nil {
		log.Fatal(err)
	}

	strptr := func(s string) *string {
		return &s
	}

	data := &model.TodoCreateInputType{
		Status:   strptr("other"),
		Priority: strptr("low"),
		Category: strptr("dummy"),
		Task:     strptr("dummy task"),
	}

	repo := repo.New(file, "dummy.json")
	repo.Create(data)
}
