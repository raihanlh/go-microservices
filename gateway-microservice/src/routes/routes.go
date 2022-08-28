package routes

import (
	"context"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
	pb "github.com/raihanlh/gateway-microservice/proto"
	"github.com/raihanlh/gateway-microservice/src/middlewares"
)

type Router interface {
	Route(app *fiber.App)
}

func RouteAll(app *fiber.App, authService pb.AuthServiceClient, routers ...Router) {
	// Use authentication middleware
	app.Use(middlewares.NewAuthMiddleware(middlewares.AuthMiddlewareConfig{
		Filter: func(c *fiber.Ctx) bool {
			path := c.OriginalURL()
			method := c.Method()

			if contains(path, method, protected) {
				token := c.Get("Authorization", "")
				if strings.HasPrefix(token, "Bearer ") {
					token = strings.Split(token, "Bearer ")[1]
				}

				req := &pb.GetByTokenRequest{
					Token: token,
				}

				_, err := authService.GetByToken(context.Background(), req)
				return err == nil
			}

			return true

		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusUnauthorized)
		},
	}))

	for _, router := range routers {
		router.Route(app)
	}

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Not found")
	})
}

func contains(path string, method string, protecteds []HTTPRequest) bool {
	for _, protected := range protecteds {
		matched_method := method == protected.Method
		if matched, _ := regexp.MatchString(protected.Path, path); matched && matched_method {
			return true
		}
	}
	return false
}
