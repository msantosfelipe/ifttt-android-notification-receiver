package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	API_PREFIX    string
	PORT          string
	VALID_API_KEY string
	VALID_APPS    []string
}

type Email struct {
	ENABLE_EMAIL   bool
	EMAIL_SERVER   string
	EMAIL_PORT     int
	EMAIL_FROM     string
	EMAIL_TO       string
	EMAIL_USERNAME string
	EMAIL_PASSWORD string
}

var ENV Config
var EMAIL_ENV Email

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
		VALID_APPS:    parseList(os.Getenv("VALID_APPS")),
	}

	EMAIL_ENV = Email{
		ENABLE_EMAIL:   parseBool(os.Getenv("ENABLE_EMAIL")),
		EMAIL_SERVER:   os.Getenv("EMAIL_SERVER"),
		EMAIL_PORT:     parseInt(os.Getenv("EMAIL_PORT")),
		EMAIL_FROM:     os.Getenv("EMAIL_FROM"),
		EMAIL_TO:       os.Getenv("EMAIL_TO"),
		EMAIL_USERNAME: os.Getenv("EMAIL_USERNAME"),
		EMAIL_PASSWORD: os.Getenv("EMAIL_PASSWORD"),
	}
}

func parseList(value string) []string {
	return strings.Split(value, ",")
}

func parseBool(value string) bool {
	return value == "true"
}

func parseInt(value string) int {
	intValue, erro := strconv.Atoi(value)
	if erro != nil {
		log.Fatalf("Error parsing int value")
	}
	return intValue
}
