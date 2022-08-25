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

	authAddress := fmt.Sprintf("%v:%v", "", configuration.Auth.Port)

	var authConn *grpc.ClientConn
	authConn, err = grpc.Dial(authAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to port %v: %v", configuration.Auth.Port, err)
	}

	defer authConn.Close()

	authService := pb.NewAuthServiceClient(authConn)
	authRouter := routes.NewAuthRouter(authService)

	app := fiber.New()

	routes.RouteAll(app, authService, authRouter)
	log.Fatal(app.Listen(":3000"))
}
