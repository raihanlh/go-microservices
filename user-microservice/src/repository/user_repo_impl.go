package repository

import (
	"context"
	"database/sql"
	"time"

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
	const query = `INSERT INTO user_details (id_user, fullname, id_gender, phone, date_of_birth, created_at, updated_at, deleted_at) ` +
		`VALUES (?, ?, ?, ?, ?, ?, ?, ?) RETURNING id, fullname, id_gender, phone, date_of_birth, created_at, updated_at`

	var id int64
	var fullname string
	var id_gender string
	var phone string
	var dob time.Time
	var created_at string
	var updated_at string

	loc := time.FixedZone("UTC+7", 7*60*60)
	dob = time.Date(int(user.DateOfBirth.Year), time.Month(int(user.DateOfBirth.Month)), int(user.DateOfBirth.Day), 0, 0, 0, 0, loc)

	stmt, err := repo.DB.PrepareContext(c, query)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRowContext(c, user.IdAccount, user.Fullname, user.IdGender, user.Phone, dob, user.CreatedAt.AsTime(), user.UpdatedAt.AsTime(), nil)

	err = row.Scan(c, &fullname, &id_gender, &phone, &dob, &created_at, &updated_at)
	if err != nil {
		return nil, err
	}

	user.Id = id

	return user, nil
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
