package repo

import (
	"errors"
	"time"

	"github.com/masputrawae/todo-cli/internal/model"
	"github.com/masputrawae/todo-cli/internal/utils"
)

var (
	ErrNotChanged   = errors.New("Nothing has changed")
	ErrIDNotFound   = errors.New("id not found")
	ErrTodoNotFound = errors.New("todo not found")
)

type TodoRepo struct {
	Data     []model.Todo
	FilePath string
}

type TodoManage interface {
	Add(u model.TodoAddInput) error
	Edit(u model.TodoEditInput) error
	Delete(id int) error
	FindAll() ([]model.Todo, error)
	FindByID(id int) ([]model.Todo, error)
	FindByStatus(status string) ([]model.Todo, error)
	FindByPriority(priority string) ([]model.Todo, error)
	FindByProject(project string) ([]model.Todo, error)
}

func NewTodoRepo(data []model.Todo) TodoManage {
	return &TodoRepo{Data: data}
}

// add new todo
func (tr *TodoRepo) Add(u model.TodoAddInput) error {
	var todo model.Todo

	todo.ID = utils.GenTodoID(tr.Data)
	todo.Task = u.Task

	if u.Priority != nil {
		todo.Priority = u.Priority
	}

	if u.Status != nil {
		todo.Status = u.Status
	}

	if u.Project != nil {
		todo.Project = u.Project
	}

	tr.Data = append(tr.Data, todo)
	return utils.SaveTodo(tr.FilePath, tr.Data)
}

// edit todo
func (tr *TodoRepo) Edit(u model.TodoEditInput) error {
	updated := false
	index := -1
	for i := range tr.Data {
		if tr.Data[i].ID == u.ID {
			index = i
		}
	}
	if index != -1 {
		if u.Task != nil {
			tr.Data[index].Task = *u.Task
			updated = true
		}
		if u.Priority != nil {
			tr.Data[index].Priority = u.Priority
			updated = true
		}
		if u.Status != nil {
			tr.Data[index].Status = u.Status
			updated = true
		}
		if u.Project != nil {
			tr.Data[index].Status = u.Status
			updated = true
		}
		if updated {
			dt := time.Now()
			tr.Data[index].Lastmod = &dt
		}

		return utils.SaveTodo(tr.FilePath, tr.Data)
	}

	return ErrNotChanged
}

// delete todo
func (tr *TodoRepo) Delete(id int) error {
	index := -1
	for i := range tr.Data {
		if tr.Data[i].ID == id {
			index = i
		}
	}
	if index != -1 {
		tr.Data = append(tr.Data[:index], tr.Data[index+1:]...)
		return utils.SaveTodo(tr.FilePath, tr.Data)
	}
	return ErrIDNotFound
}

// find all todos
func (tr *TodoRepo) FindAll() ([]model.Todo, error) {
	var results = make([]model.Todo, len(tr.Data))
	copy(results, tr.Data)
	if len(results) == 0 {
		return results, ErrTodoNotFound
	}
	return results, nil
}

// find todo by id
func (tr *TodoRepo) FindByID(id int) ([]model.Todo, error) {
	var results []model.Todo
	for i := range tr.Data {
		if tr.Data[i].ID == id {
			results = append(results, tr.Data[i])
		}
	}
	if len(results) == 0 {
		return results, ErrTodoNotFound
	}

	return results, nil
}

// find todo by status
func (tr *TodoRepo) FindByStatus(status string) ([]model.Todo, error) {
	var results []model.Todo
	for i := range tr.Data {
		if tr.Data[i].Status == &status {
			results = append(results, tr.Data[i])
		}
	}
	if len(results) == 0 {
		return results, ErrTodoNotFound
	}
	return results, nil
}

// find todo by Priority
func (tr *TodoRepo) FindByPriority(priority string) ([]model.Todo, error) {
	var results []model.Todo
	for i := range tr.Data {
		if tr.Data[i].Priority == &priority {
			results = append(results, tr.Data[i])
		}
	}
	if len(results) == 0 {
		return results, ErrTodoNotFound
	}
	return results, nil
}

// find todo by project
func (tr *TodoRepo) FindByProject(project string) ([]model.Todo, error) {
	var results []model.Todo
	for i := range tr.Data {
		if tr.Data[i].Project == &project {
			results = append(results, tr.Data[i])
		}
	}
	if len(results) == 0 {
		return results, ErrTodoNotFound
	}
	return results, nil
}
