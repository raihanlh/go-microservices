package routes

import (
	"context"
	"fmt"
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
			// var unprotected = []string{"\\/login", "\\/register", "\\/articles\\/[\\d]"}
			var protected = []string{"\\/user"}

			// if contains(path, unprotected) {
			// 	return true
			// }

			if contains(path, protected) {
				token := c.Get("Authorization", "")
				fmt.Println(token)
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

func contains(el string, s []string) bool {
	for _, a := range s {
		if matched, _ := regexp.MatchString(a, el); matched {
			return true
		}
		// if a == el {
		// 	return true
		// }
	}
	fmt.Println(("FALSE"))
	return false
}
