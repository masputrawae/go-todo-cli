package utils

import (
	"sort"

	"github.com/masputrawae/todo-cli/internal/model"
)

func GenTodoID(t []model.Todo) int {
	sort.Slice(t, func(i, j int) bool { return t[i].ID < t[j].ID })
	id := 1
	for i := range t {
		if t[i].ID == id {
			id++
		}
	}
	return id
}
