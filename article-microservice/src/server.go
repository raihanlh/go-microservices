package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	_ "github.com/lib/pq"

	pb "github.com/raihanlh/go-article-microservice/proto"
	"github.com/raihanlh/go-article-microservice/src/config"
	"github.com/raihanlh/go-article-microservice/src/repository"
	"github.com/raihanlh/go-article-microservice/src/service"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Running gRPC article server...")

	var db *sql.DB

	configuration, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configuration.DB.Host, configuration.DB.Port, configuration.DB.User, configuration.DB.Password, configuration.DB.Name)

	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Connect to auth service
	authAddress := fmt.Sprintf("%v:%v", "", configuration.Auth.Port)

	var authConn *grpc.ClientConn
	authConn, err = grpc.Dial(authAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to port %v: %v", configuration.Auth.Port, err)
	}
	fmt.Println("Connected succesfully")

	defer authConn.Close()

	authService := pb.NewAuthServiceClient(authConn)

	// Listen to port
	port := fmt.Sprintf(":%v", configuration.Article.Port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen to port %v: %v", configuration.Article.Port, err)
	}

	articleRepository := repository.NewArticleRepository(db)

	articleServer := service.ArticleServer{
		ArticleRepository: articleRepository,
		AuthService:       authService,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterArticleServiceServer(grpcServer, &articleServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %v: %v", configuration.Auth.Port, err)
	}
}
