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
	pb "github.com/raihanlh/go-article-microservice/proto"
	"github.com/raihanlh/go-article-microservice/src/config"
	"github.com/raihanlh/go-article-microservice/src/repository"
	"github.com/raihanlh/go-article-microservice/src/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// Create a metrics registry.
	reg = prometheus.NewRegistry()

	// Create some standard server metrics.
	grpcMetrics = grpc_prometheus.NewServerMetrics()

	// Create a customized counter metric.
	getArticleCounterMetric = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "get_article_count",
		Help: "Total number of article fetched on the server.",
	}, []string{"article_id", "title"})
)

func init() {
	// Register standard server metrics and customized metrics to registry.
	reg.MustRegister(grpcMetrics, getArticleCounterMetric)
}

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
	authAddress := fmt.Sprintf("%v:%v", configuration.Auth.Host, configuration.Auth.Port)

	var authConn *grpc.ClientConn
	authConn, err = grpc.Dial(authAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
		ArticleCounterVec: getArticleCounterMetric,
	}

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpcMetrics.StreamServerInterceptor()),
		grpc.UnaryInterceptor(grpcMetrics.UnaryServerInterceptor()),
	)

	pb.RegisterArticleServiceServer(grpcServer, &articleServer)

	grpcMetrics.InitializeMetrics(grpcServer)

	httpServer := &http.Server{Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}), Addr: fmt.Sprintf(":%d", 3102)}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal("Unable to start a http server.")
		}
	}()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %v: %v", configuration.Auth.Port, err)
	}
}
