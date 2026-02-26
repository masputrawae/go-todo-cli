package service

import (
	"errors"

	"github.com/masputrawae/go-todo-cli/pkg/model"
)

func (s *Service) GetTodo(i *model.TodoFindType) ([]model.Todo, error) {
	results := s.Repo.Find(i)
	if len(results) == 0 {
		return nil, errors.New("tugas tidak ditemukan")
	}
	return results, nil
}
