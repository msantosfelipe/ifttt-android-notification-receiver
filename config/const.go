package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/msantosfelipe/ifttt-android-notification-receiver/domain"
)

var ENV domain.Config
var EMAIL_ENV domain.Email
var PUSH_NOTIFICATION_ENV domain.PushNotification

func InitVars() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	ENV = domain.Config{
		API_PREFIX:     os.Getenv("API_PREFIX"),
		PORT:           os.Getenv("PORT"),
		VALID_API_KEY:  os.Getenv("VALID_API_KEY"),
		ALLOW_ALL_APPS: parseBool(os.Getenv("ALLOW_ANY_APP")),
		ALLOWED_APPS:   parseList(os.Getenv("ALLOWED_APPS")),
		ALLOWED_TITLES: parseList(os.Getenv("ALLOWED_TITLES")),
	}

	EMAIL_ENV = domain.Email{
		ENABLE:         parseBool(os.Getenv("ENABLE_EMAIL")),
		EMAIL_SERVER:   os.Getenv("EMAIL_SERVER"),
		EMAIL_PORT:     parseInt(os.Getenv("EMAIL_PORT")),
		EMAIL_FROM:     os.Getenv("EMAIL_FROM"),
		EMAIL_TO:       os.Getenv("EMAIL_TO"),
		EMAIL_USERNAME: os.Getenv("EMAIL_USERNAME"),
		EMAIL_PASSWORD: os.Getenv("EMAIL_PASSWORD"),
	}

	PUSH_NOTIFICATION_ENV = domain.PushNotification{
		ENABLE_ONE_SIGNAL:       parseBool(os.Getenv("ENABLE_ONE_SIGNAL_NOTIFICATION")),
		ENABLE_PUSHOVER:         parseBool(os.Getenv("ENABLE_PUSHOVER_NOTIFICATION")),
		ONE_SIGNAL_APP_ID:       os.Getenv("ONE_SIGNAL_APP_ID"),
		ONE_SIGNAL_REST_API_KEY: os.Getenv("ONE_SIGNAL_REST_API_KEY"),
		ONE_SIGNAL_SEGMENT:      os.Getenv("ONE_SIGNAL_SEGMENT"),
		PUSH_OVER_APP_TOKEN:     os.Getenv("PUSH_OVER_APP_TOKEN"),
		PUSH_OVER_APP_RECIPIENT: os.Getenv("PUSH_OVER_APP_RECIPIENT"),
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
