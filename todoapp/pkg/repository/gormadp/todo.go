package gormadp

import (
	"gorm.io/gorm"
	"todoapp/pkg/repository/gormadp/dbmodels"
	"todoapp/pkg/repository/models"
)

var _ Interface = (*TodoRepository)(nil)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (t *TodoRepository) ListTodos() ([]*models.Todo, error) {
	var dbTodos []dbmodels.Todo

	if err := t.db.Find(&dbTodos).Error; err != nil {
		return nil, err
	}

	todoList := make([]*models.Todo, 0, len(dbTodos))

	for i := range dbTodos {
		todoList = append(todoList, dbTodos[i].To())
	}

	return todoList, nil
}

func (t *TodoRepository) AddTodo(m *models.Todo) error {
	var dbTodo dbmodels.Todo

	dbTodo.From(m)

	if err := t.db.Create(&dbTodo).Error; err != nil {
		return err
	}

	return nil
}

func (t *TodoRepository) UpdateTodo(id int, done bool) error {
	if err := t.db.Model(&dbmodels.Todo{}).Where("id = ?", id).UpdateColumn("done", done).Error; err != nil {
		return err
	}

	return nil
}

func (t *TodoRepository) DeleteTodo(id int) error {
	if err := t.db.Delete(&dbmodels.Todo{}, id).Error; err != nil {
		return err
	}

	return nil
}
