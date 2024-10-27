package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func NotificationHandler(c *fiber.Ctx) error {
	var body any
	if err := c.BodyParser(&body); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request",
		})
	}

	fmt.Println(body)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Notification received",
	})
}
