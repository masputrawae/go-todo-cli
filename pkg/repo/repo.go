package repo

import (
	"sort"

	"github.com/masputrawae/todo-cli/pkg/model"
	"github.com/masputrawae/todo-cli/pkg/utils"
)

type Repo struct {
	TodoFile string
	Todos    []model.Todo
}

type Manage interface {
	Create(input *model.TodoCreateInputType) error
}

func New(todo []model.Todo, todoFile string) Manage {
	return &Repo{TodoFile: todoFile, Todos: todo}
}

// ===== TODOS MANAGE =====
func (r *Repo) Create(input *model.TodoCreateInputType) error {
	sort.Slice(r.Todos, func(i, j int) bool { return r.Todos[i].ID < r.Todos[j].ID })
	id := 1

	for i := range r.Todos {
		if r.Todos[i].ID == id {
			id++
		}
	}

	r.Todos = append(r.Todos, model.Todo{
		ID:       id,
		Status:   *input.Status,
		Priority: *input.Priority,
		Category: *input.Category,
		Task:     *input.Task,
	})

	return utils.SaveData(r.TodoFile, r.Todos)
}
