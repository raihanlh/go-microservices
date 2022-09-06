package v1

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	pb "github.com/raihanlh/gateway-microservice/proto"
	"github.com/raihanlh/gateway-microservice/src/routes"
	"github.com/raihanlh/gateway-microservice/src/util"
	"google.golang.org/grpc/status"
)

type AuthRouter struct {
	AuthService pb.AuthServiceClient
}

func NewAuthRouter(authService pb.AuthServiceClient) routes.Router {
	return &AuthRouter{
		AuthService: authService,
	}
}

func (a *AuthRouter) Route(app *fiber.App) {
	app.Post("/register", a.Register)
	app.Post("/login", a.Login)
	app.Get("/user", a.GetByToken)
	app.Get("/hello", a.Hello)
}

func (a *AuthRouter) Hello(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"success": true,
		"data":    "hello",
	})
}

func (a *AuthRouter) Register(ctx *fiber.Ctx) error {
	var req pb.RegisterRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	res, err := a.AuthService.Register(context.Background(), &req)
	if err != nil {
		log.Println("failed to register", err.Error())
		if e, ok := status.FromError(err); ok {
			return ctx.Status(util.HTTPStatusFromCode(e.Code())).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		} else {
			return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"success": true,
		"data":    res,
	})
}

func (a *AuthRouter) Login(ctx *fiber.Ctx) error {
	var req pb.LoginRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	res, err := a.AuthService.Login(context.Background(), &req)
	if err != nil {
		log.Println("failed to login", err.Error())
		if e, ok := status.FromError(err); ok {
			return ctx.Status(util.HTTPStatusFromCode(e.Code())).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		} else {
			return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"success": true,
		"data":    res,
	})
}

func (a *AuthRouter) GetByToken(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	token := strings.Split(authHeader, "Bearer ")[1]

	req := &pb.GetByTokenRequest{
		Token: token,
	}

	account, err := a.AuthService.GetByToken(context.Background(), req)

	if err != nil {
		if e, ok := status.FromError(err); ok {
			return ctx.Status(util.HTTPStatusFromCode(e.Code())).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		} else {
			return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"success": true,
		"data": map[string]interface{}{
			"id":         account.Id,
			"email":      account.Email,
			"role":       account.Role,
			"enable":     account.Enable,
			"locked":     account.Locked,
			"created_at": account.CreatedAt.AsTime(),
			"updated_at": account.UpdatedAt.AsTime(),
		},
	})
}
