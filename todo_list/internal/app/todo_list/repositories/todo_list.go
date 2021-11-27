package repositories

import (
	"github.com/youchann/golang-practice/todo_list/internal/app/todo_list/schemas"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) GetAll() ([]schemas.Todo, error) {
	var todos []schemas.Todo
	result := r.db.Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}
	return todos, nil
}
