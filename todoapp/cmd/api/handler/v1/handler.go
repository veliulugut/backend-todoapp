package v1

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"todoapp/pkg/repository/models"
	"todoapp/service/todo"
)

var _ Interface = (*Handler)(nil)

type Handler struct {
	todoService todo.Interface
}

func New(s todo.Interface) *Handler {
	return &Handler{
		todoService: s,
	}
}

func (h *Handler) ListTodos(w http.ResponseWriter, r *http.Request) {
	var (
		todos []*models.Todo
		err   error
	)

	if todos, err = h.todoService.ListTodos(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func (h *Handler) AddTodo(w http.ResponseWriter, r *http.Request) {
	var (
		icBody models.Todo
		err    error
	)

	if err = json.NewDecoder(r.Body).Decode(&icBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = h.todoService.AddTodo(&icBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var (
		icBody models.Todo
		err    error
	)

	if err = json.NewDecoder(r.Body).Decode(&icBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idParam := chi.URLParam(r, "id")

	var id int64

	if id, err = strconv.ParseInt(idParam, 10, 64); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = h.todoService.UpdateTodo(int(id), icBody.Done); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	var (
		id  int64
		err error
	)

	if id, err = strconv.ParseInt(idParam, 10, 64); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = h.todoService.DeleteTodo(int(id)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
