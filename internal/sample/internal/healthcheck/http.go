package healthcheck

import (
	"github.com/gofiber/fiber/v2"
)

// RegisterHandlers register handler for health check endpoint
func RegisterHandlers(r fiber.Router) {
	r.Get("/health", Check)
}

// Check implements endpoint for health check
func Check(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "OK",
		"error":   false,
	})
}
