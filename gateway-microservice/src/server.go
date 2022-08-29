package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	pb "github.com/raihanlh/gateway-microservice/proto"
	"github.com/raihanlh/gateway-microservice/src/config"
	"github.com/raihanlh/gateway-microservice/src/routes"
	routes_v1 "github.com/raihanlh/gateway-microservice/src/routes/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	configuration, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Connect to auth microservice
	authAddress := fmt.Sprintf("%v:%v", configuration.Auth.Host, configuration.Auth.Port)

	var authConn *grpc.ClientConn
	authConn, err = grpc.Dial(authAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to port %v: %v", configuration.Auth.Port, err)
	}

	defer authConn.Close()

	authService := pb.NewAuthServiceClient(authConn)
	authRouter := routes_v1.NewAuthRouter(authService)

	// Connect to article microservice
	articleAddress := fmt.Sprintf("%v:%v", configuration.Auth.Host, configuration.Article.Port)
	var articleConn *grpc.ClientConn
	articleConn, err = grpc.Dial(articleAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to port %v: %v", configuration.Article.Port, err)
	}

	defer articleConn.Close()

	articleService := pb.NewArticleServiceClient(articleConn)
	articleRouter := routes_v1.NewArticleRouter(articleService)

	// Create fiber
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		},
	})

	routes.RouteAll(app, authService, authRouter, articleRouter)
	log.Fatal(app.Listen(":3000"))
}
