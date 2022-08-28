package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raihanlh/gateway-microservice/src/routes"
)

type UserRouter struct{}

func NewUserRouter() routes.Router {
	return &UserRouter{}
}

func (a *UserRouter) Route(app *fiber.App) {
	// var conn *grpc.ClientConn
	// conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	// if err != nil {
	// 	log.Fatalf("Could not connect to port 9000: %v", err)
	// }

	// defer conn.Close()

	// app.Get("/users", getBooks(service))
	// app.Post("/users", addBook(service))
	// app.Put("/users/:id", updateBook(service))
	// app.Delete("/books", removeBook(service))
}
