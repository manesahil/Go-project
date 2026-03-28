package config

import (
	"log"
	"os"

	"Project/Model"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	dsn := os.Getenv("DB_CONN_STR")
	if dsn == "" {
		log.Fatal("DB_CONN_STR not set in .env")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db

	// Auto migrate models
	err = db.AutoMigrate(&Model.User{})
	if err != nil {
		log.Fatal("Failed to auto-migrate models:", err)
	}

	log.Println("Database connected and migrated successfully!")
}
