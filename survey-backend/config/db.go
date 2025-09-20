package config

import (
	"fmt"
	"log"
	"os"

	"github.com/monikalilhare01/survey-backend/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() (*gorm.DB, error) {
	_ = godotenv.Load()

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	params := "parseTime=true&charset=utf8mb4&loc=Local"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", user, pass, host, port, name, params)
	fmt.Println("Connecting to DB with DSN:", dsn)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}

	if err := DB.AutoMigrate(
		&models.User{},
		&models.Survey{},
		&models.Questions{},
		&models.Response{},
		&models.ResponseAnswer{},
	); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Println("Database connected and migrated successfully")
	return DB, nil
}
