package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var ENV Config

type Config struct {
	API_PREFIX    string
	PORT          string
	VALID_API_KEY string
}

func InitVars() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	ENV = Config{
		API_PREFIX:    os.Getenv("API_PREFIX"),
		PORT:          os.Getenv("PORT"),
		VALID_API_KEY: os.Getenv("VALID_API_KEY"),
	}
}
