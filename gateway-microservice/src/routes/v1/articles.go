package v1

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	pb "github.com/raihanlh/gateway-microservice/proto"
	"github.com/raihanlh/gateway-microservice/src/routes"
	"github.com/raihanlh/gateway-microservice/src/util"
	"google.golang.org/grpc/status"
)

type ArticleRouter struct {
	ArticleService pb.ArticleServiceClient
}

func NewArticleRouter(articleService pb.ArticleServiceClient) routes.Router {
	return &ArticleRouter{
		ArticleService: articleService,
	}
}

func (a *ArticleRouter) Route(app *fiber.App) {
	app.Get("/articles/:id", a.GetArticle)
	app.Post("/articles", a.CreateArticle)
	// app.Put("/articles/:id", a.UpdateArticle)
	// app.Delete("/articles", DeleteArticle)
}

func (a *ArticleRouter) GetArticle(ctx *fiber.Ctx) error {
	id, _ := strconv.ParseInt(ctx.Params("id"), 10, 64)
	res, err := a.ArticleService.GetArticleById(context.Background(), &pb.GetArticleRequest{
		Id: id,
	})
	if err != nil {
		log.Printf("failed to get article with id %d: %s\n", id, err.Error())
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

func (a *ArticleRouter) CreateArticle(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	token := strings.Split(authHeader, "Bearer ")[1]

	var req pb.CreateArticleRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	res, err := a.ArticleService.CreateArticle(context.Background(), &pb.CreateArticleRequest{
		Title:   req.Title,
		Content: req.Content,
		Token:   token,
	})
	if err != nil {
		log.Printf("failed create article: %s\n", err.Error())
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

// func (a *ArticleRouter) UpdateArticle(ctx *fiber.Ctx) error {
// 	var req pb.UpdateArticleRequest
// 	err := ctx.BodyParser(&req)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return err
// 	}
// 	fmt.Println(&req)
// 	return nil
// }
