package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"testing"

	_ "github.com/lib/pq"
	pb "github.com/raihanlh/go-user-microservice/proto"
	"github.com/raihanlh/go-user-microservice/src/config"
	"github.com/raihanlh/go-user-microservice/src/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
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
		"localhost", configuration.DB.User, configuration.DB.Password, configuration.DB.Name)

	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}

	// defer db.Close()

	// Connect to auth service
	authAddress := fmt.Sprintf("%v:%v", "localhost", configuration.Auth.Port)

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
	userDetailRepository := repository.NewUserDetailRepository(db)
	pb.RegisterUserDetailServiceServer(s, &UserServer{
		UserDetailRepository: userDetailRepository,
		AuthService:          authService,
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

func TestUserDetailService(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserDetailServiceClient(conn)
	auth_token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAbG9jYWwuaG9zdCIsImV4cCI6MTY2Njc5NzI2MiwiaWQiOjF9.XsKEU5gz1hWmgZfoSTQErnKOyGYse7HcVW7yolrBuqc"

	// t.Run("Ensure create/update user detail is success", func(t *testing.T) {
	// 	res, err := client.CreateUpdateUserDetail(ctx, &pb.CreateUpdateUserDetailRequest{
	// 		Fullname: "Test McTester",
	// 		IdGender: 0,
	// 		Phone:    "08123456789",
	// 		DateOfBirth: &pb.Date{
	// 			Day:   1,
	// 			Month: 1,
	// 			Year:  1997,
	// 		},
	// 		Token: auth_token,
	// 	})
	// 	if err != nil {
	// 		t.Log(err)
	// 		t.Errorf("Create/update user detail test failed")
	// 	}
	// 	PrintResult(res)
	// })

	t.Run("Ensure get user detail is success", func(t *testing.T) {
		res, err := client.GetUserDetailByUser(ctx, &pb.GetUserDetailByUserRequest{
			Token: auth_token,
		})
		if err != nil {
			t.Log(err)
			t.Errorf("Create/update user detail test failed")
		}
		PrintResult(res)
	})
}

func PrintResult(res interface{}) error {
	result, err := json.MarshalIndent(res, "", " ")
	if err != nil {
		return err
	}
	fmt.Println(string(result))
	return nil
}
