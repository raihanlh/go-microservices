package repository

import (
	"database/sql"
	"time"

	"github.com/raihanlh/go-auth-microservice/src/entity"
)

type AccountRepositoryImpl struct {
	DB *sql.DB
}

func NewAccountRepository(db *sql.DB) AccountRepository {
	return &AccountRepositoryImpl{
		DB: db,
	}
}

func (repo *AccountRepositoryImpl) Save(account *entity.Account) (int64, error) {
	// Prepare statement
	const query = `INSERT INTO accounts (username, email, password, id_role, enable) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id int64

	// Query to db and return id
	err := repo.DB.QueryRow(query, account.Email, account.Email, account.Password, 2, true).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (repo *AccountRepositoryImpl) FindAll(size int, page int, sortBy string, sortType string) (*[]entity.Account, error) {
	// const query = `SELECT * FROM accounts`
	// var accounts *[]entity.Account

	// err := repo.DB.Query(query).Scan(&accounts)
	// if err != nil {
	// 	return nil, err
	// }

	// return accounts, nil
	return nil, nil
}

func (repo *AccountRepositoryImpl) FindById(id int64) (entity.Account, error) {
	const query = `SELECT email, id_role, enable, locked, created_at, updated_at FROM accounts a WHERE a.id = $1`
	var email string
	var role int64
	var enable bool
	var locked bool
	var created_at time.Time
	var updated_at time.Time

	err := repo.DB.QueryRow(query, id).Scan(&email, &role, &enable, &locked, &created_at, &updated_at)

	if err != nil {
		return entity.Account{}, err
	}

	return entity.Account{
		Id:        id,
		Email:     email,
		Role:      role,
		Enable:    enable,
		Locked:    locked,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
	}, nil
}

func (repo *AccountRepositoryImpl) FindByEmail(email string) (entity.Account, error) {
	const query = `SELECT id, password FROM accounts a WHERE a.email = $1`
	var id int64
	var password string

	err := repo.DB.QueryRow(query, email).Scan(
		&id, &password)

	if err != nil {
		return entity.Account{}, err
	}

	return entity.Account{
		Id:       id,
		Email:    email,
		Password: password,
	}, nil
}

func (repo *AccountRepositoryImpl) FindPasswordByEmail(email string) (string, error) {
	const query = `SELECT password FROM accounts a WHERE a.email = $1`
	var password string

	err := repo.DB.QueryRow(query, email).Scan(&password)

	if err != nil {
		return "", err
	}

	return password, nil
}

func (repo *AccountRepositoryImpl) Update(account *entity.Account) (*entity.Account, error) {
	return nil, nil
}

func (repo *AccountRepositoryImpl) Delete(account *entity.Account) error {
	return nil
}

func (repo *AccountRepositoryImpl) IsExist(email string) (bool, error) {
	const query = `SELECT id FROM accounts a WHERE a.email = $1`
	var id int64
	err := repo.DB.QueryRow(query, email).Scan(&id)

	if err != nil && err != sql.ErrNoRows {
		return false, err
	} else if err == sql.ErrNoRows {
		return false, nil
	} else {
		return true, nil
	}
}
