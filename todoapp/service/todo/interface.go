package todo

import "todoapp/pkg/repository/models"

type Interface interface {
	ListTodos() ([]*models.Todo, error)
	AddTodo(t *models.Todo) error
	UpdateTodo(id int, done bool) error
	DeleteTodo(id int) error
}
