package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/msantosfelipe/ifttt-android-notification-receiver/config"
	"github.com/msantosfelipe/ifttt-android-notification-receiver/handler"
	"github.com/msantosfelipe/ifttt-android-notification-receiver/handler/middleware"
	"github.com/msantosfelipe/ifttt-android-notification-receiver/infra"
	"github.com/msantosfelipe/ifttt-android-notification-receiver/usecase"
)

func main() {
	config.InitVars()

	logEnabledSenders()

	// Initialize dependencies
	mailSender := infra.NewMailSender()
	pushNotificationSender := infra.NewPushNotificationSender()
	uc := usecase.NewNotificationUsecase(mailSender, pushNotificationSender)
	notificationHandler := handler.NewNotificationHandler(uc)

	// Init server
	startServer(notificationHandler)
}

func logEnabledSenders() {
	if config.EMAIL_ENV.ENABLE {
		fmt.Println("Email is enabled")
	}

	if config.PUSH_NOTIFICATION_ENV.ENABLE {
		fmt.Println("Push notification is enabled")
	}
}

func startServer(notificationHandler handler.NotificationHandler) {
	fmt.Println("Starting server...")
	app := fiber.New()

	app.Use(middleware.ApiKeyMiddleware)
	app.Post(fmt.Sprintf("%s/receive", config.ENV.API_PREFIX), notificationHandler.ProcessNotification)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.ENV.PORT)))
}
