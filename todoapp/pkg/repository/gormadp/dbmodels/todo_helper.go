package dbmodels

import "todoapp/pkg/repository/models"

func (t *Todo) From(m *models.Todo) {
	*t = Todo{
		ID:   m.ID,
		Text: m.Text,
		Done: m.Done,
	}
}

func (t *Todo) To() *models.Todo {
	return &models.Todo{
		ID:   t.ID,
		Text: t.Text,
		Done: t.Done,
	}
}
