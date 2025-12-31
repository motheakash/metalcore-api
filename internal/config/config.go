package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Try to load from current directory first
	err := godotenv.Load()

	// If not found, try parent directory (for when running from tmp/)
	if err != nil {
		err = godotenv.Load("../.env")
	}

	if err != nil {
		log.Println("No .env file found.")
	}
}
