package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("postgres database cannot load .env file")
	}

	var conErr error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	DB, conErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if conErr != nil {
		log.Fatal("postgres database connection fail", dsn)
	}

	log.Println("postgres database connected")
}
