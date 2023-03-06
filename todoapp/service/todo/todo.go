package todo

import (
	"todoapp/pkg/repository/gormadp"
	"todoapp/pkg/repository/models"
)

var _ Interface = (*Service)(nil)

type Service struct {
	repo *gormadp.Repository
}

func New(repo *gormadp.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ListTodos() ([]*models.Todo, error) {
	var (
		todos []*models.Todo
		err   error
	)

	if todos, err = s.repo.TodoRepo.ListTodos(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (s *Service) AddTodo(t *models.Todo) error {
	if err := s.repo.TodoRepo.AddTodo(t); err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateTodo(id int, done bool) error {
	if err := s.repo.TodoRepo.UpdateTodo(id, done); err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteTodo(id int) error {
	if err := s.repo.TodoRepo.DeleteTodo(id); err != nil {
		return err
	}

	return nil
}
