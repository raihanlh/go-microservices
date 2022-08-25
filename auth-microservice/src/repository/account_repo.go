package repository

import "github.com/raihanlh/go-auth-microservice/src/entity"

type AccountRepository interface {
	Save(account *entity.Account) (int64, error)
	FindAll(size int, page int, sortBy string, sortType string) (*[]entity.Account, error)
	FindById(id int64) (entity.Account, error)
	FindByEmail(email string) (entity.Account, error)
	FindPasswordByEmail(email string) (string, error)
	Update(account *entity.Account) (*entity.Account, error)
	Delete(account *entity.Account) error
	IsExist(email string) (bool, error)
}
