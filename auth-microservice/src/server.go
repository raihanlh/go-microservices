package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	pb "github.com/raihanlh/go-auth-microservice/proto"
	"github.com/raihanlh/go-auth-microservice/src/config"
	"github.com/raihanlh/go-auth-microservice/src/repository"
	"github.com/raihanlh/go-auth-microservice/src/service"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Running gRPC auth server...")

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

	port := fmt.Sprintf(":%v", configuration.Auth.Port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen to port %v: %v", configuration.Auth.Port, err)
	}

	accountRepository := repository.NewAccountRepository(db)

	authServer := service.AuthServer{
		AccountRepository: accountRepository,
	}

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)

	pb.RegisterAuthServiceServer(grpcServer, &authServer)

	grpc_prometheus.Register(grpcServer)

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(":3101", nil))
	}()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %v: %v", configuration.Auth.Port, err)
	}
}
