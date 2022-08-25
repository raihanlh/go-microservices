package routes

import "github.com/gofiber/fiber/v2"

type ArticleRouter struct{}

func NewArticleRouter() Router {
	return &ArticleRouter{}
}

func (a *ArticleRouter) Route(app *fiber.App) {
	// app.Get("/books", getBooks(service))
	// app.Post("/books", addBook(service))
	// app.Put("/books", updateBook(service))
	// app.Delete("/books", removeBook(service))
}
