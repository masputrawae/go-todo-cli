package service

import (
	"context"
	"strings"

	"github.com/masputrawae/todo-cli/internal/data"
	"github.com/masputrawae/todo-cli/internal/model"
	"github.com/masputrawae/todo-cli/internal/repo"
)

type TodoAPI struct {
	Todos      []model.Todo
	Statuses   []model.Meta
	Priorities []model.Meta
}

type Service struct {
	Repo       repo.Manage
	Statauses  []model.Meta
	Priorities []model.Meta
}

type Manage interface {
	Create(ctx context.Context, input *repo.CreateInputType) error
	Update(ctx context.Context, input *repo.UpdateInputType) error
	Delete(ctx context.Context, input *repo.DeleteInputType) error
	GetTodos(ctx context.Context, input *repo.FindInputType) (TodoAPI, error)
}

func New(r repo.Manage, statuses, priorities []model.Meta) Manage {
	return &Service{Repo: r, Statauses: statuses, Priorities: priorities}
}

// utils for checking meta (statuses & categories)
func metaMathces(m []model.Meta, i string) bool {
	normalize := func(s string) string {
		return strings.ReplaceAll(strings.ToLower(s), " ", "-")
	}
	for _, v := range m {
		if normalize(v.Name) == normalize(i) {
			return true
		}
	}
	return false
}

func (s *Service) Create(ctx context.Context, input *repo.CreateInputType) error {
	todo := input
	todo.Priority = &data.Default.Prioriy
	todo.Status = &data.Default.Status
	if input.Priority != nil {
		if metaMathces(s.Priorities, *input.Priority) {
			todo.Priority = input.Priority
		}
	}
	if input.Status != nil {
		if metaMathces(s.Statauses, *input.Status) {
			todo.Status = input.Status
		}
	}

	return s.Repo.Create(ctx, todo)
}

func (s *Service) Update(ctx context.Context, input *repo.UpdateInputType) error {
	todo := input
	if input.Priority != nil {
		if !metaMathces(s.Priorities, *input.Priority) {
			todo.Priority = &data.Default.Prioriy
		}
	}
	if input.Status != nil {
		if !metaMathces(s.Statauses, *input.Status) {
			todo.Status = &data.Default.Status
		}
	}

	return s.Repo.Update(ctx, todo)
}

func (s *Service) Delete(ctx context.Context, input *repo.DeleteInputType) error {
	return s.Repo.Delete(ctx, input)
}

func (s *Service) GetTodos(ctx context.Context, input *repo.FindInputType) (TodoAPI, error) {
	todos, err := s.Repo.Find(ctx, input)
	if err != nil {
		return TodoAPI{}, err
	}

	return TodoAPI{
		Todos:      todos,
		Statuses:   s.Statauses,
		Priorities: s.Priorities,
	}, nil
}
