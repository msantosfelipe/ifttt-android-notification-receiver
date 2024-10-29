package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var ENV Config
var EMAIL_ENV Email
var PUSH_NOTIFICATION_ENV PushNotification

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
		ENABLE:         parseBool(os.Getenv("ENABLE_EMAIL")),
		EMAIL_SERVER:   os.Getenv("EMAIL_SERVER"),
		EMAIL_PORT:     parseInt(os.Getenv("EMAIL_PORT")),
		EMAIL_FROM:     os.Getenv("EMAIL_FROM"),
		EMAIL_TO:       os.Getenv("EMAIL_TO"),
		EMAIL_USERNAME: os.Getenv("EMAIL_USERNAME"),
		EMAIL_PASSWORD: os.Getenv("EMAIL_PASSWORD"),
	}

	PUSH_NOTIFICATION_ENV = PushNotification{
		ENABLE:                  parseBool(os.Getenv("ENABLE_PUSH_NOTIFICATION")),
		ONE_SIGNAL_APP_ID:       os.Getenv("ONE_SIGNAL_APP_ID"),
		ONE_SIGNAL_REST_API_KEY: os.Getenv("ONE_SIGNAL_REST_API_KEY"),
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
