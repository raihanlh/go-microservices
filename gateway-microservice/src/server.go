package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	pb "github.com/raihanlh/gateway-microservice/proto"
	"github.com/raihanlh/gateway-microservice/src/config"
	"github.com/raihanlh/gateway-microservice/src/routes"
	"google.golang.org/grpc"
)

func main() {
	configuration, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Connect to auth microservice
	authAddress := fmt.Sprintf("%v:%v", "", configuration.Auth.Port)

	var authConn *grpc.ClientConn
	authConn, err = grpc.Dial(authAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to port %v: %v", configuration.Auth.Port, err)
	}

	defer authConn.Close()

	authService := pb.NewAuthServiceClient(authConn)
	authRouter := routes.NewAuthRouter(authService)

	// Connect to article microservice
	articleAddress := fmt.Sprintf("%v:%v", "", configuration.Article.Port)
	var articleConn *grpc.ClientConn
	articleConn, err = grpc.Dial(articleAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to port %v: %v", configuration.Article.Port, err)
	}

	defer articleConn.Close()

	articleService := pb.NewArticleServiceClient(articleConn)
	articleRouter := routes.NewArticleRouter(articleService)

	// Create fiber
	app := fiber.New()

	routes.RouteAll(app, authService, authRouter, articleRouter)
	log.Fatal(app.Listen(":3000"))
}
