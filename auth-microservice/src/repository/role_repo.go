package repository

import "github.com/raihanlh/go-auth-microservice/src/entity"

type RoleRepository interface {
	FindOneByName(name string) (*entity.Role, error)
}
