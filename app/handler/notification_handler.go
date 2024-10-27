package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func NotificationHandler(c *fiber.Ctx) error {
	fmt.Printf("Method: %s\n", c.Method())
	fmt.Printf("URL: %s\n", c.OriginalURL())
	fmt.Printf("Headers: %v\n", c.GetReqHeaders())
	fmt.Printf("Body: %s\n", string(c.Body()))

	var body any

	if err := c.BodyParser(&body); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Notification received",
	})
}
