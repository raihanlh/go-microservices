package service

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	pb "github.com/raihanlh/go-auth-microservice/proto"
	"github.com/raihanlh/go-auth-microservice/src/config"
	"github.com/raihanlh/go-auth-microservice/src/repository"
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

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	accountRepository := repository.NewAccountRepository(db)
	pb.RegisterAuthServiceServer(s, &AuthServer{
		AccountRepository: accountRepository,
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

func TestAuthService(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewAuthServiceClient(conn)

	t.Run("Ensure login is success", func(t *testing.T) {
		res, err := client.Login(ctx, &pb.LoginRequest{
			Email:    "test@local.host",
			Password: "password",
		})
		if err != nil {
			t.Errorf("Login test failed")
		}

		t.Log(res)
		fmt.Println(res)
	})

	t.Run("Ensure get by id is success", func(t *testing.T) {
		res, err := client.GetByToken(ctx, &pb.GetByTokenRequest{
			Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAbG9jYWwuaG9zdCIsImV4cCI6MTY2MDk2OTg3NiwiaWQiOjl9._H6depccI9b3u0vy8vr3aHmr-r2Ybp06Q1xBcOm_OBo",
		})
		if err != nil {
			t.Errorf("Get by token test failed: %s", err.Error())
		}

		t.Log(res)
		// fmt.Println(res)
	})
}
