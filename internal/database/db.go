// internal/database/db.go
package database

import (
	"tinyrun/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("tinyrun.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&models.User{}, &models.Run{})
	return db
}
