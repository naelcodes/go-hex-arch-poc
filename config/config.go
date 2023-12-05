package config

import (
	"flag"
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

	stage := flag.String("stage", "dev", "stage")

	flag.Parse()

	// Print stage
	log.Printf("Stage: %s", *stage)

	// Load environment variables
	if *stage == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file", err)
		}
	}

	POSTGRES_DB_HOST = os.Getenv("POSTGRES_DB_HOST")
	POSTGRES_DB_USER = os.Getenv("POSTGRES_DB_USER")
	POSTGRES_DB_PASSWORD = os.Getenv("POSTGRES_DB_PASSWORD")
	POSTGRES_DB_NAME = os.Getenv("POSTGRES_DB_NAME")
	POSTGRES_DB_PORT = os.Getenv("POSTGRES_DB_PORT")
	APP_ENGINE_SERVER_PORT = os.Getenv("APP_ENGINE_SERVER_PORT")

	// Print environment variables
	// log.Printf("POSTGRES_DB_HOST: %s", POSTGRES_DB_HOST)
	// log.Printf("POSTGRES_DB_USER: %s", POSTGRES_DB_USER)
	// log.Printf("POSTGRES_DB_PASSWORD: %s", POSTGRES_DB_PASSWORD)
	// log.Printf("POSTGRES_DB_NAME: %s", POSTGRES_DB_NAME)
	// log.Printf("POSTGRES_DB_PORT: %s", POSTGRES_DB_PORT)
	// log.Printf("APP_ENGINE_SERVER_PORT: %s", APP_ENGINE_SERVER_PORT)

}
