package infrastructures

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDbConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("../../../../todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
