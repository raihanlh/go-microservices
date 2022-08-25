package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

type AuthMiddlewareConfig struct {
	Filter       func(c *fiber.Ctx) bool // Required
	Unauthorized fiber.Handler           // middleware specfic
}

func NewAuthMiddleware(config AuthMiddlewareConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if config.Filter != nil && config.Filter(c) {
			return c.Next()
		}
		return config.Unauthorized(c)
	}
}
