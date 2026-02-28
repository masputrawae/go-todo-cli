package service

import (
	"errors"

	"github.com/masputrawae/todo-cli/internal/model"
	"github.com/masputrawae/todo-cli/internal/repo"
	"github.com/masputrawae/todo-cli/internal/utils"
)

var (
	ErrPriorityNotSame = errors.New("priorities are not the same")
)

type TodoService struct {
	Repo       repo.TodoManage
	Statuses   []model.Status
	Priorities []model.Priority
}

type TodoServiceManage interface {
	Create(u model.TodoAddInput) error
}

func NewTodoService(r repo.TodoManage) TodoServiceManage {
	return &TodoService{Repo: r}
}

func (ts *TodoService) Create(u model.TodoAddInput) error {
	var todo model.TodoAddInput
	todo.Task = u.Task

	if u.Priority != nil {
		for i := range ts.Priorities {
			sameID := utils.TextIsSame(ts.Priorities[i].ID, *u.Priority)
			sameShort := utils.TextIsSame(ts.Priorities[i].Short, *u.Priority)
			if !sameID || sameShort {
				return ErrPriorityNotSame
			}
			priority := utils.NormalizeText(*u.Priority)
			todo.Priority = &priority
		}
	}

	if u.Status != nil {
		for i := range ts.Statuses {
			if !utils.TextIsSame(ts.Statuses[i].ID, *u.Status) {
				return ErrPriorityNotSame
			}
			status := utils.NormalizeText(*u.Status)
			todo.Priority = &status
		}
	}
	return ts.Repo.Add(todo)
}
