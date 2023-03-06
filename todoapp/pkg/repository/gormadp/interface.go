package gormadp

import "todoapp/pkg/repository/models"

type Interface interface {
	ListTodos() ([]*models.Todo, error)
	AddTodo(m *models.Todo) error
	UpdateTodo(id int, done bool) error
	DeleteTodo(id int) error
}
