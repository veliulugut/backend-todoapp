package v1

import "net/http"

type Interface interface {
	ListTodos(w http.ResponseWriter, r *http.Request)
	AddTodo(w http.ResponseWriter, r *http.Request)
	UpdateTodo(w http.ResponseWriter, r *http.Request)
	DeleteTodo(w http.ResponseWriter, r *http.Request)
}
