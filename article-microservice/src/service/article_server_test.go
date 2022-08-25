package service

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	pb "github.com/raihanlh/go-article-microservice/proto"
	"github.com/raihanlh/go-article-microservice/src/config"
	"github.com/raihanlh/go-article-microservice/src/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	var db *sql.DB

	configuration, err := config.LoadConfigByPath("../../..")
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		configuration.DB.Host, configuration.DB.User, configuration.DB.Password, configuration.DB.Name)

	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}

	// defer db.Close()

	// Connect to auth service
	authAddress := fmt.Sprintf("%v:%v", "", configuration.Auth.Port)

	var authConn *grpc.ClientConn
	authConn, err = grpc.Dial(authAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to port %v: %v", configuration.Auth.Port, err)
	}
	fmt.Println("Connected succesfully")

	// defer authConn.Close()

	authService := pb.NewAuthServiceClient(authConn)

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	articleRepository := repository.NewArticleRepository(db)
	pb.RegisterArticleServiceServer(s, &ArticleServer{
		ArticleRepository: articleRepository,
		AuthService:       authService,
	})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestArticleService(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewArticleServiceClient(conn)

	t.Run("Ensure create article is success", func(t *testing.T) {
		res, err := client.CreateArticle(ctx, &pb.CreateArticleRequest{
			Title:   "Test",
			Content: "Test content",
			Token:   "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAbG9jYWwuaG9zdCIsImV4cCI6MTY2MTUwMDUyMCwiaWQiOjl9.O63PMfYfxAhf54_4ANuFc4vlgL4yjPSNNyZdcuZb6fE",
		})
		if err != nil {
			t.Log(err)
			t.Errorf("Create article test failed")
		}
		t.Log(res)
		fmt.Println(res)
	})
}
