package database

import (
	"fmt"
	"log"
	"os"

	"github.com/bipro-b/ecommerce-backend/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env file not found or skipped")
	}

	// Use TEST_DB_NAME if GO_ENV=test
	dbName := os.Getenv("DB_NAME")
	if os.Getenv("GO_ENV") == "test" {
		dbName = os.Getenv("TEST_DB_NAME")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		dbName,
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	// Run AutoMigrate for all models
	if err := db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}); err != nil {
		log.Fatalf("❌ AutoMigrate failed: %v", err)
	}

	DB = db
	log.Println("✅ Successfully connected to database")
}
