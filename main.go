package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/msantosfelipe/ifttt-android-notification-receiver/app/config"
	"github.com/msantosfelipe/ifttt-android-notification-receiver/app/handler"
	"github.com/msantosfelipe/ifttt-android-notification-receiver/app/handler/middleware"
)

func main() {
	config.InitVars()

	fmt.Println("Starting server...")
	app := fiber.New()

	app.Use(middleware.ApiKeyMiddleware)

	app.Post(fmt.Sprintf("%s/receive", config.ENV.API_PREFIX), handler.NotificationHandler)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.ENV.PORT)))
}
