package service

import (
	pb "github.com/raihanlh/go-user-microservice/proto"
	"github.com/raihanlh/go-user-microservice/src/repository"
)

type UserService struct {
	pb.UnimplementedUserDetailServiceServer
	UserRepository repository.UserRepository
	AuthService    pb.AuthServiceClient
}
