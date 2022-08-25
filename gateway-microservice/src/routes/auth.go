package routes

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	pb "github.com/raihanlh/gateway-microservice/proto"
	"github.com/raihanlh/gateway-microservice/src/util"
	"google.golang.org/grpc/status"
)

type AuthRouter struct {
	AuthService pb.AuthServiceClient
}

func NewAuthRouter(authService pb.AuthServiceClient) Router {
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
	log.Println(&req)
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
	log.Println(&req)
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
			log.Println(e)
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

type AccountDTO struct {
	Id        int64     `json:"id"`
	Email     string    `json:"email"`
	Role      int64     `json:"role"`
	Enable    bool      `json:"enable"`
	Locked    bool      `json:"locked"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (a *AuthRouter) GetByToken(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	token := strings.Split(authHeader, "Bearer ")[1]

	req := &pb.GetByTokenRequest{
		Token: token,
	}

	account, err := a.AuthService.GetByToken(context.Background(), req)
	res := &AccountDTO{
		Id:        account.Id,
		Email:     account.Email,
		Role:      account.Role,
		Enable:    account.Enable,
		Locked:    account.Locked,
		CreatedAt: account.CreatedAt.AsTime(),
		UpdatedAt: account.UpdatedAt.AsTime(),
	}

	if err != nil {
		if e, ok := status.FromError(err); ok {
			log.Println(e)
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
