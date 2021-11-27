package repositories

import (
	"github.com/youchann/golang-practice/todo_list/internal/app/todo_list/schemas/models"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	result := r.db.Find(&todos)
	if result.Error != nil {
		return nil, result.Error
	}
	return todos, nil
}

func (r *TodoRepository) Create(todo models.Todo) (models.Todo, error) {
	result := r.db.Create(&todo)
	if result.Error != nil {
		return todo, result.Error
	}
	return todo, nil
}
