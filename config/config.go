package config

import (
	"fmt"
	"go-backend/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectDB() {
	LoadEnv()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Run migrations for all models
	err = DB.AutoMigrate(&models.User{}, &models.Product{}) // ✅ Add Product model
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	fmt.Println("Database migration completed successfully")
}

func CloseDB() {
	db, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}
