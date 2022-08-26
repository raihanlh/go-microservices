package service

import (
	"context"
	"database/sql"
	// "encoding/json"
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
	authConn, err = grpc.Dial(authAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to port %v: %v", configuration.Auth.Port, err)
	}
	fmt.Println("Connected to auth gRPC server succesfully")

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
	var id int64
	auth_token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAbG9jYWwuaG9zdCIsImV4cCI6MTY2MTc3NDMyMCwiaWQiOjl9._YicKqEmN6M7NY8ZZkaVk6N0B8e9UOUvl_7UavS7NQ4"

	t.Run("Ensure create article is success", func(t *testing.T) {
		res, err := client.CreateArticle(ctx, &pb.CreateArticleRequest{
			Title:   "Test",
			Content: "Test content",
			Token:   auth_token,
		})
		if err != nil {
			t.Log(err)
			t.Errorf("Create article test failed")
		}
		t.Log(res)
		id = res.Article.Id
	})

	t.Run("Ensure get article is success", func(t *testing.T) {
		res, err := client.GetArticleById(ctx, &pb.GetArticleRequest{
			Id: id,
		})
		if err != nil {
			t.Log(err)
			t.Errorf("Get article test failed")
		}
		t.Log(res)
	})

	t.Run("Ensure get all article owned by a user id is success", func(t *testing.T) {
		res, err := client.GetArticleByUser(ctx, &pb.GetAllArticleByUserRequest{
			Token: auth_token,
		})
		if err != nil {
			t.Log(err)
			t.Errorf("Get article by user token test failed")
		}
		t.Log(res)

		// result, err := json.MarshalIndent(res, "", " ")
		// if err != nil {
		// 	t.Log(err)
		// }
		// fmt.Println(string(result))
	})

	t.Run("Ensure get all article is success", func(t *testing.T) {
		res, err := client.GetAllArticle(ctx, &pb.GetAllArticleRequest{})
		if err != nil {
			t.Log(err)
			t.Errorf("Get all article test failed")
		}
		t.Log(res)

		// result, err := json.MarshalIndent(res, "", " ")
		// if err != nil {
		// 	t.Log(err)
		// }
		// fmt.Println(string(result))
	})
}
