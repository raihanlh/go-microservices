package service

import (
	"context"
	"log"

	pb "github.com/raihanlh/go-user-microservice/proto"
	"github.com/raihanlh/go-user-microservice/src/repository"
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
		log.Println(err)
		return nil, err
	}

	isExist, err := server.UserDetailRepository.IsExist(ctx, user.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if !isExist {
		res, err := server.UserDetailRepository.Save(ctx, &pb.UserDetail{})
		log.Println(err)
		if err != nil {
			return nil, err
		}
		return res, nil
	} else {
		res, err := server.UserDetailRepository.Update(ctx, &pb.UserDetail{})
		log.Println(err)
		if err != nil {
			return nil, err
		}
		return res, nil
	}

}
