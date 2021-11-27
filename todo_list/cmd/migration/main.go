package main

import (
	"github.com/youchann/golang-practice/todo_list/internal/app/todo_list/schemas/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("../../todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&models.Todo{})
	db.Create(&models.Todo{Name: "hoge", IsDone: false})
}
