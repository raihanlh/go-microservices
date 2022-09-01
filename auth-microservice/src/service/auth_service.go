package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/prometheus/client_golang/prometheus"
	pb "github.com/raihanlh/go-auth-microservice/proto"
	"github.com/raihanlh/go-auth-microservice/src/config"
	"github.com/raihanlh/go-auth-microservice/src/entity"
	"github.com/raihanlh/go-auth-microservice/src/repository"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	AccountRepository repository.AccountRepository
	CounterVec        *prometheus.CounterVec
}

func (server *AuthServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	log.Println("registering ", req.Email)

	isExist, err := server.AccountRepository.IsExist(req.Email)
	if err != nil {
		log.Println("register failed: ", err)
		return nil, err
	}
	if isExist {
		log.Println("register failed: email already exist")
		return nil, status.Error(codes.AlreadyExists, "email already exist")
	}

	if req.Password != req.ConfirmPassword {
		log.Println("register failed: password and confirm password doesn't match")
		return &pb.RegisterResponse{}, status.Error(codes.InvalidArgument, "password and confirm password doesn't match")
	}

	passwordHashed, err := hashPassword(req.Password)
	if err != nil {
		log.Println("failed to hash password: ", err.Error())
		return nil, err
	}

	account := &entity.Account{
		Email:    req.Email,
		Password: passwordHashed,
	}

	accountId, err := server.AccountRepository.Save(account)
	if err != nil {
		log.Println("failed to save account: ", err.Error())
		return nil, err
	}

	t, err := generateToken(req.Email, accountId)
	if err != nil {
		log.Println("failed to create token: ", err.Error())
		return nil, err
	}

	log.Println("register success")
	return &pb.RegisterResponse{
		Token: t,
	}, nil
}

func (server *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Println("logging in: ", req.Email)

	account, err := server.AccountRepository.FindByEmail(req.Email)

	if err != nil {
		log.Println("failed to create token: ", err.Error())
		return nil, status.Error(codes.NotFound, "email not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(req.Password))
	if err != nil {
		log.Println("incorrect password: ", err.Error())
		return nil, status.Error(codes.InvalidArgument, "incorrect password")
	}

	token, err := generateToken(req.Email, account.Id)
	if err != nil {
		log.Println("failed to create token", err.Error())
		return nil, err
	}

	server.CounterVec.WithLabelValues(req.Email).Inc()
	return &pb.LoginResponse{
		Token: token,
	}, nil
}

func (server *AuthServer) GetByToken(ctx context.Context, req *pb.GetByTokenRequest) (*pb.GetUserResponse, error) {
	fmt.Println((req))
	configuration, err := config.LoadConfig()
	if err != nil {
		log.Println("failed to load configuration", err.Error())
		return &pb.GetUserResponse{}, err
	}

	var signingKey = []byte(configuration.JWT.Secret)

	if req.Token == "" {
		log.Println("no token")
		return nil, status.Error(codes.InvalidArgument, "no token")
	}

	token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return signingKey, nil
	})

	if err != nil {
		log.Println("Your Token has been expired: ", err.Error())
		return nil, status.Error(codes.InvalidArgument, "your token has been expired")
	}

	// If token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		account, err := server.AccountRepository.FindById(int64(claims["id"].(float64)))
		if err != nil {
			return &pb.GetUserResponse{}, status.Error(codes.NotFound, err.Error())
		}
		return &pb.GetUserResponse{
			Id:        account.Id,
			Email:     account.Email,
			Role:      account.Role,
			Enable:    account.Enable,
			Locked:    account.Locked,
			CreatedAt: timestamppb.New(account.CreatedAt),
			UpdatedAt: timestamppb.New(account.UpdatedAt),
		}, nil
	}

	return nil, status.Error(codes.InvalidArgument, "Invalid token")
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func generateToken(email string, id int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = email
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	configuration, err := config.LoadConfig()
	if err != nil {
		log.Println("failed to load configuration", err.Error())
		return "", err
	}

	t, err := token.SignedString([]byte(configuration.JWT.Secret))
	if err != nil {
		log.Println("failed to create token", err.Error())
		return "", err
	}

	return t, nil
}
