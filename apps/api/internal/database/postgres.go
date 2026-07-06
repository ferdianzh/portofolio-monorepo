package database

import (
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
	dsn := os.Getenv("DATABASE_URL")

	DB, conErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if conErr != nil {
		log.Fatal("postgres database connection fail", dsn)
	}

	log.Println("postgres database connected")
}
