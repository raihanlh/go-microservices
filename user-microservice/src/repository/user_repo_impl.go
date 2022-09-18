package repository

import (
	"context"
	"database/sql"
	"time"

	pb "github.com/raihanlh/go-user-microservice/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var loc = time.FixedZone("UTC+7", 7*60*60)

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
	var id_gender int64
	var phone string
	var dob time.Time
	var created_at time.Time
	var updated_at time.Time

	dob = time.Date(int(user.DateOfBirth.Year), time.Month(int(user.DateOfBirth.Month)), int(user.DateOfBirth.Day), 0, 0, 0, 0, loc)

	stmt, err := repo.DB.PrepareContext(c, query)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRowContext(c, user.IdAccount, user.Fullname, user.IdGender, user.Phone, dob, user.CreatedAt.AsTime(), user.UpdatedAt.AsTime(), nil)

	err = row.Scan(&id, &fullname, &id_gender, &phone, &dob, &created_at, &updated_at)
	if err != nil {
		return nil, err
	}

	user.Id = id
	user.CreatedAt = timestamppb.New(created_at)
	user.UpdatedAt = timestamppb.New(updated_at)

	return user, nil
}

func (repo *UserRepositoryImpl) Update(c context.Context, user *pb.UserDetail) (*pb.UserDetail, error) {
	const query = `UPDATE user_details u SET fullname = ?, id_gender = ?, phone = ?, date_of_birth = ?, updated_at = ? WHERE id_user = ? AND deleted_at IS NULL RETURNING created_at, updated_at`

	var created_at time.Time
	var updated_at time.Time

	dob := time.Date(int(user.DateOfBirth.Year), time.Month(int(user.DateOfBirth.Month)), int(user.DateOfBirth.Day), 0, 0, 0, 0, loc)

	stmt, err := repo.DB.PrepareContext(c, query)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRowContext(c, user.Fullname, user.IdGender, user.Phone, dob, user.UpdatedAt.AsTime().In(loc), user.IdAccount)
	err = row.Scan(&created_at, &updated_at)
	if err != nil {
		return nil, err
	}

	user.CreatedAt = timestamppb.New(created_at)
	user.UpdatedAt = timestamppb.New(updated_at)

	return user, nil
}

func (repo *UserRepositoryImpl) FindByAccountId(c context.Context, id_account int64) (*pb.UserDetail, error) {
	query := `SELECT id, id_user, fullname, id_gender, phone, date_of_birth, created_at, updated_at, deleted_at FROM user_details WHERE id_user = ? AND deleted_at IS NULL`

	var id int64
	var id_user int64
	var fullname string
	var id_gender int64
	var phone string
	var dob time.Time
	var created_at time.Time
	var updated_at time.Time
	var deleted_at sql.NullTime

	stmt, err := repo.DB.PrepareContext(c, query)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRowContext(c, id_account)
	err = row.Scan(&id, &id_user, &fullname, &id_gender, &phone, &dob, &created_at, &updated_at, &deleted_at)
	if err != nil {
		return nil, err
	}

	year, month, day := time.Now().Date()

	return &pb.UserDetail{
		Id:        id,
		IdAccount: id_user,
		Fullname:  fullname,
		IdGender:  id_gender,
		Phone:     phone,
		DateOfBirth: &pb.Date{
			Day:   int32(day),
			Month: int32(month),
			Year:  int32(year),
		},
		CreatedAt: timestamppb.New(created_at),
		UpdatedAt: timestamppb.New(updated_at),
	}, nil
}

func (repo *UserRepositoryImpl) FindAll(c context.Context) ([]*pb.UserDetail, error) {
	const query = `SELECT id, id_user, fullname, id_gender, phone, date_of_birth, created_at, updated_at, deleted_at FROM user_details WHERE deleted_at IS NULL`
	stmt, err := repo.DB.PrepareContext(c, query)
	if err != nil {
		return nil, err
	}

	user_details := make([]*pb.UserDetail, 0)
	rows, err := stmt.QueryContext(c, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id int64
		var id_user int64
		var fullname string
		var id_gender int64
		var phone string
		var dob time.Time
		var created_at time.Time
		var updated_at time.Time
		var deleted_at sql.NullTime

		err = rows.Scan(&id, &id_user, &fullname, &id_gender, &phone, &dob, &created_at, &updated_at, &deleted_at)
		if err != nil {
			return nil, err
		}
		year, month, day := time.Now().Date()

		user_details = append(user_details, &pb.UserDetail{
			Id:        id,
			IdAccount: id_user,
			Fullname:  fullname,
			IdGender:  id_gender,
			Phone:     phone,
			DateOfBirth: &pb.Date{
				Day:   int32(day),
				Month: int32(month),
				Year:  int32(year),
			},
			CreatedAt: timestamppb.New(created_at),
			UpdatedAt: timestamppb.New(updated_at),
		})

	}

	return user_details, nil
}

func (repo *UserRepositoryImpl) IsExist(c context.Context, id_account int64) (bool, error) {
	const query = `SELECT id FROM user_details u WHERE u.id_user = ?`
	var id int64

	stmt, err := repo.DB.PrepareContext(c, query)
	if err != nil {
		return false, err
	}

	err = stmt.QueryRowContext(c, id_account).Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	} else if err == sql.ErrNoRows {
		return false, nil
	} else {
		return true, nil
	}
}
