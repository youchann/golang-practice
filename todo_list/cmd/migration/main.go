package main

import (
	"github.com/youchann/golang-practice/todo_list/internal/app/todo_list/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("../../todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&schemas.Todo{})
	db.Create(&schemas.Todo{Name: "hoge", IsDone: false})
}
