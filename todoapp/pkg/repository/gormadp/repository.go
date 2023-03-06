package gormadp

import (
	"gorm.io/gorm"
	"todoapp/pkg/repository/gormadp/dbmodels"
)

type Repository struct {
	TodoRepo *TodoRepository
	db       *gorm.DB
}

func New(db *gorm.DB) *Repository {
	todoRepo := NewTodoRepository(db)

	return &Repository{TodoRepo: todoRepo, db: db}
}

func (r *Repository) Migrate() error {
	if err := r.db.AutoMigrate(
		&dbmodels.Todo{},
	); err != nil {
		return err
	}

	return nil
}
