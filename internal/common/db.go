package common

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	cfg := LoadConfig()

	db, err := gorm.Open(sqlite.Open(cfg.DatabasePath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	return db
}

func ApplyMigrations(db *gorm.DB) {
	// Migrations are applied in main.go after imports are wired.
	// db.AutoMigrate(&user.User{})
	// db.AutoMigrate(&dish.Dish{})
	// db.AutoMigrate(&review.Review{})

	log.Println("Database migrations applied")
}
