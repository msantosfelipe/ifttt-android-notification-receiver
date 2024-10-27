package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/msantosfelipe/ifttt-android-notification-receiver/domain"
)

type notificationHandler struct {
	notificationUc domain.NotificationUsecase
}

type NotificationHandler interface {
	ProcessNotification(c *fiber.Ctx) error
}

func NewNotificationHandler(notificationUc domain.NotificationUsecase) NotificationHandler {
	return &notificationHandler{
		notificationUc: notificationUc,
	}
}

func (handler *notificationHandler) ProcessNotification(c *fiber.Ctx) error {
	fmt.Println("Notification received.")
	fmt.Printf("Body: %s\n", string(c.Body()))

	var notification domain.Notification

	if err := c.BodyParser(&notification); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request",
		})
	}

	handler.notificationUc.ProcessNotification(notification)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Notification received",
	})
}
