package user

import (
	"github.com/gofiber/fiber/v2"
)

// RegisterHttpHandlers sets up the routing of the HTTP handlers
func RegisterHttpHandlers(r fiber.Router, h ApiHttpHandler) {
	r.Get("/", h.GetUserList)
	r.Get("/:id", h.GetUserByID)
}

// RegisterAdminHttpHandlers register non auth api
func RegisterAdminHttpHandlers(r fiber.Router, h NoAuthHttpHandler) {
	//
}

// RegisterNoAuthHttpHandlers register non auth api
func RegisterNoAuthHttpHandlers(r fiber.Router, h AdminHttpHandler) {
	//
}
