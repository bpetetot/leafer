package db

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // Use Sqlite dialect
)

// Setup migrates all models and returns a DB connection
func Setup() *gorm.DB {
	database := os.Getenv("DATABASE")

	db, err := gorm.Open("sqlite3", database)

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Library{})
	db.AutoMigrate(&Media{})

	return db
}
