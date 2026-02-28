package service

import (
	"errors"

	"github.com/masputrawae/todo-cli/internal/model"
	"github.com/masputrawae/todo-cli/internal/repo"
	"github.com/masputrawae/todo-cli/internal/utils"
)

var (
	ErrPriorityNotSame = errors.New("priorities are not the same")
	ErrStatusNotSame   = errors.New("statuses are not the same")
)

type TodoService struct {
	Repo       repo.TodoManage
	Statuses   []model.Status
	Priorities []model.Priority
}

type TodoServiceManage interface {
	Load() error
	Create(u model.TodoAddInput) error
	Update(u model.TodoEditInput) error
	Delete(id int) error

	GetAll() ([]model.Todo, error)
	GetByID(id int) ([]model.Todo, error)
	GetByStatus(status string) ([]model.Todo, error)
	GetByPriority(priority string) ([]model.Todo, error)
	GetByProject(project string) ([]model.Todo, error)
}

func NewTodoService(r repo.TodoManage, st []model.Status, pr []model.Priority) TodoServiceManage {
	return &TodoService{Repo: r, Statuses: st, Priorities: pr}
}

func (ts *TodoService) Load() error {
	return ts.Repo.Load()
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
				return ErrStatusNotSame
			}
			status := utils.NormalizeText(*u.Status)
			todo.Priority = &status
		}
	}

	if u.Project != nil {
		project := utils.NormalizeText(*u.Project)
		todo.Project = &project
	}

	return ts.Repo.Add(todo)
}

func (ts *TodoService) Update(u model.TodoEditInput) error {
	todos, err := ts.Repo.FindByID(u.ID)
	if err != nil {
		return err
	}

	old := todos[0]
	var todo model.TodoEditInput
	todo.ID = u.ID

	if u.Task != nil && utils.TextIsSame(old.Task, *u.Task) {
		todo.Task = u.Task
	}

	if u.Priority != nil && utils.TextIsSame(*old.Priority, *u.Priority) {
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

	if u.Status != nil && utils.TextIsSame(*old.Status, *u.Status) {
		for i := range ts.Statuses {
			if !utils.TextIsSame(ts.Statuses[i].ID, *u.Status) {
				return ErrStatusNotSame
			}
			status := utils.NormalizeText(*u.Status)
			todo.Priority = &status
		}
	}

	if u.Project != nil && utils.TextIsSame(*old.Project, *u.Project) {
		project := utils.NormalizeText(*u.Project)
		todo.Project = &project
	}

	return ts.Repo.Edit(todo)
}

func (ts *TodoService) Delete(id int) error {
	if _, err := ts.Repo.FindByID(id); err != nil {
		return nil
	}
	return ts.Repo.Delete(id)
}

func (ts *TodoService) GetAll() ([]model.Todo, error) {
	return ts.Repo.FindAll()
}

func (ts *TodoService) GetByID(id int) ([]model.Todo, error) {
	return ts.Repo.FindByID(id)
}

func (ts *TodoService) GetByPriority(priority string) ([]model.Todo, error) {
	pr := utils.NormalizeText(priority)
	return ts.Repo.FindByPriority(pr)
}

func (ts *TodoService) GetByStatus(status string) ([]model.Todo, error) {
	st := utils.NormalizeText(status)
	return ts.Repo.FindByPriority(st)
}

func (ts *TodoService) GetByProject(project string) ([]model.Todo, error) {
	pr := utils.NormalizeText(project)
	return ts.Repo.FindByProject(pr)
}
