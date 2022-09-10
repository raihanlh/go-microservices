package repository

import (
	"context"
	"database/sql"

	pb "github.com/raihanlh/go-user-microservice/proto"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (repo *UserRepositoryImpl) Save(c context.Context, user *pb.UserDetail) (*pb.UserDetail, error) {
	return nil, nil
}

func (repo *UserRepositoryImpl) Update(c context.Context, user *pb.UserDetail) (*pb.UserDetail, error) {
	return nil, nil
}

func (repo *UserRepositoryImpl) FindByAccountId(c context.Context, id_account int64) (*pb.UserDetail, error) {
	return nil, nil
}

func (repo *UserRepositoryImpl) FindAll(c context.Context) ([]*pb.UserDetail, error) {
	return nil, nil
}
