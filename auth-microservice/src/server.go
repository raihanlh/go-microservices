package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	pb "github.com/raihanlh/go-auth-microservice/proto"
	"github.com/raihanlh/go-auth-microservice/src/config"
	"github.com/raihanlh/go-auth-microservice/src/repository"
	"github.com/raihanlh/go-auth-microservice/src/service"
	"google.golang.org/grpc"
)

var (
	// Create a metrics registry.
	reg = prometheus.NewRegistry()

	// Create some standard server metrics.
	grpcMetrics = grpc_prometheus.NewServerMetrics()

	// Create a customized counter metric.
	customizedCounterMetric = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "login_handle_count",
		Help: "Total number of Login RPCs handled on the server.",
	}, []string{"email"})
)

func init() {
	// Register standard server metrics and customized metrics to registry.
	reg.MustRegister(grpcMetrics, customizedCounterMetric)
}

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
		CounterVec:        customizedCounterMetric,
	}

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpcMetrics.StreamServerInterceptor()),
		grpc.UnaryInterceptor(grpcMetrics.UnaryServerInterceptor()),
	)

	pb.RegisterAuthServiceServer(grpcServer, &authServer)

	// grpc_prometheus.Register(grpcServer)
	grpcMetrics.InitializeMetrics(grpcServer)

	httpServer := &http.Server{Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}), Addr: fmt.Sprintf("%v:%d", configuration.Auth.Host, 3101)}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal("Unable to start a http server.")
		}
	}()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %v: %v", configuration.Auth.Port, err)
	}
}
