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
	app.Use(middlewares.NewAuthMiddleware(middlewares.AuthMiddlewareConfig{
		Filter: func(c *fiber.Ctx) bool {
			path := c.OriginalURL()
			var unprotected = []string{"\\/login", "\\/register", "\\/article\\/[\\d]"}

			if contains(unprotected, path) {
				return true
			}

			authHeader := c.Get("Authorization")
			token := strings.Split(authHeader, "Bearer ")[1]

			req := &pb.GetByTokenRequest{
				Token: token,
			}

			_, err := authService.GetByToken(context.Background(), req)
			return err == nil
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusUnauthorized)
		},
	}))

	for _, router := range routers {
		router.Route(app)
	}
}

func contains(s []string, el string) bool {
	for _, a := range s {
		if matched, _ := regexp.MatchString(a, el); matched {
			return true
		}
		// if a == el {
		// 	return true
		// }
	}
	return false
}
