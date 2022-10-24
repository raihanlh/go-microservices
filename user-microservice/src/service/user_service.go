package service

import (
	"context"
	"log"

	pb "github.com/raihanlh/go-user-microservice/proto"
	"github.com/raihanlh/go-user-microservice/src/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	pb.UnimplementedUserDetailServiceServer
	UserDetailRepository repository.UserDetailRepository
	AuthService          pb.AuthServiceClient
}

func (server *UserServer) CreateUpdateUserDetail(ctx context.Context, req *pb.CreateUpdateUserDetailRequest) (*pb.UserDetail, error) {
	authReq := &pb.GetByTokenRequest{
		Token: req.Token,
	}
	var user *pb.GetUserResponse
	user, err := server.AuthService.GetByToken(ctx, authReq)

	if err != nil {
		log.Println("Failed to get token: ", err)
		return nil, err
	}

	isExist, err := server.UserDetailRepository.IsExist(ctx, user.Id)
	log.Println(isExist)
	if err != nil {
		log.Println("Failed to get isExist: ", err)
		return nil, err
	}
	if !isExist {
		log.Println("Creating user detail...")
		res, err := server.UserDetailRepository.Save(ctx, &pb.UserDetail{
			IdAccount:   user.Id,
			Fullname:    req.Fullname,
			IdGender:    req.IdGender,
			Phone:       req.Phone,
			DateOfBirth: req.DateOfBirth,
		})
		if err != nil {
			log.Println("Failed to save user detail: ", err)
			return nil, err
		}
		return res, nil
	} else {
		res, err := server.UserDetailRepository.Update(ctx, &pb.UserDetail{})
		if err != nil {
			log.Println("Failed to update user detail: ", err)
			return nil, err
		}
		return res, nil
	}

}

func (server *UserServer) GetUserDetailByUser(ctx context.Context, req *pb.GetUserDetailByUserRequest) (*pb.UserDetail, error) {
	authReq := &pb.GetByTokenRequest{
		Token: req.Token,
	}
	var user *pb.GetUserResponse
	user, err := server.AuthService.GetByToken(ctx, authReq)

	if err != nil {
		log.Println("Failed to get token: ", err)
		return nil, err
	}

	isExist, err := server.UserDetailRepository.IsExist(ctx, user.Id)
	log.Println(isExist)
	if err != nil {
		log.Println("Failed to get isExist: ", err)
		return nil, err
	}

	if !isExist {
		return &pb.UserDetail{}, status.Error(codes.NotFound, "article not found")
	} else {
		log.Println("Getting user detail...")
		res, err := server.UserDetailRepository.FindByAccountId(ctx, user.Id)
		if err != nil {
			log.Println("Failed to get user detail: ", err)
			return nil, err
		}
		return res, nil
	}
}
