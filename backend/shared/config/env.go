package config

import (
	"github.com/joho/godotenv"
	"log"
)

// LoadEnv loads environment variables from a .env file if it exists.
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found, using environment variables: %v", err)
	}
}
