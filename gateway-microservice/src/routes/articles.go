package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	pb "github.com/raihanlh/gateway-microservice/proto"
)

type ArticleRouter struct {
	ArticleService pb.ArticleServiceClient
}

func NewArticleRouter(articleService pb.ArticleServiceClient) Router {
	return &ArticleRouter{
		ArticleService: articleService,
	}
}

func (a *ArticleRouter) Route(app *fiber.App) {
	// app.Get("/articles", GetArticle)
	// app.Post("/articles", CreateArticle)
	app.Put("/articles/:id", a.UpdateArticle)
	// app.Delete("/articles", DeleteArticle)
}

func (a *ArticleRouter) UpdateArticle(ctx *fiber.Ctx) error {
	var req pb.UpdateArticleRequest
	err := ctx.BodyParser(&req)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(&req)
	return nil
}
