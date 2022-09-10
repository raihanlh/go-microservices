package repository

import (
	"context"

	pb "github.com/raihanlh/go-user-microservice/proto"
)

type UserRepository interface {
	Save(c context.Context, user *pb.UserDetail) (*pb.UserDetail, error)
	Update(c context.Context, user *pb.UserDetail) (*pb.UserDetail, error)
	FindByAccountId(c context.Context, id_account int64) (*pb.UserDetail, error)
	FindAll(c context.Context) ([]*pb.UserDetail, error)
}
