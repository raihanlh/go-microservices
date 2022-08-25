package repository

import (
	"database/sql"

	"github.com/raihanlh/go-auth-microservice/src/entity"
)

type RoleRepositoryImpl struct {
	DB *sql.DB
}

func NewRoleRepository(db *sql.DB) RoleRepository {
	return &RoleRepositoryImpl{
		DB: db,
	}
}

func (repository *RoleRepositoryImpl) FindOneByName(name string) (*entity.Role, error) {
	return nil, nil
}
