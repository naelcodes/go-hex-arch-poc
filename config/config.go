package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	POSTGRES_DB_HOST       string
	POSTGRES_DB_USER       string
	POSTGRES_DB_PASSWORD   string
	POSTGRES_DB_NAME       string
	POSTGRES_DB_PORT       string
	APP_ENGINE_SERVER_PORT string
)

func LoadEnvironmentConfig() {

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	POSTGRES_DB_HOST = os.Getenv("POSTGRES_DB_HOST")
	POSTGRES_DB_USER = os.Getenv("POSTGRES_DB_USER")
	POSTGRES_DB_PASSWORD = os.Getenv("POSTGRES_DB_PASSWORD")
	POSTGRES_DB_NAME = os.Getenv("POSTGRES_DB_NAME")
	POSTGRES_DB_PORT = os.Getenv("POSTGRES_DB_PORT")
	APP_ENGINE_SERVER_PORT = os.Getenv("APP_ENGINE_SERVER_PORT")

}
