package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/msantosfelipe/ifttt-android-notification-receiver/config"
)

func ApiKeyMiddleware(c *fiber.Ctx) error {
	apikey := c.Get("apikey")
	if apikey == "" || apikey != config.ENV.VALID_API_KEY {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	}
	return c.Next()
}
