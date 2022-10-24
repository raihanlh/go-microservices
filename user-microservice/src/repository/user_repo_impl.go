package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

	pb "github.com/raihanlh/go-user-microservice/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var loc = time.FixedZone("UTC+7", 7*60*60)

type UserDetailRepositoryImpl struct {
	DB *sql.DB
}

func NewUserDetailRepository(db *sql.DB) UserDetailRepository {
	return &UserDetailRepositoryImpl{
		DB: db,
	}
}

func (repo *UserDetailRepositoryImpl) Save(c context.Context, user *pb.UserDetail) (*pb.UserDetail, error) {
	const query = `INSERT INTO users_details (id_user, fullname, id_gender, phone, date_of_birth, created_at, updated_at, deleted_at) ` +
		`VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, fullname, id_gender, phone, date_of_birth, created_at, updated_at`

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
		log.Println("Failed to prepare context for saving user detail: ", err)
		return nil, err
	}
	row := stmt.QueryRowContext(c, user.IdAccount, user.Fullname, user.IdGender, user.Phone, dob, time.Now(), time.Now(), nil)

	err = row.Scan(&id, &fullname, &id_gender, &phone, &dob, &created_at, &updated_at)
	if err != nil {
		log.Println("Failed to scan row user detail: ", err)
		return nil, err
	}

	user.Id = id
	user.CreatedAt = timestamppb.New(created_at)
	user.UpdatedAt = timestamppb.New(updated_at)

	return user, nil
}

func (repo *UserDetailRepositoryImpl) Update(c context.Context, user *pb.UserDetail) (*pb.UserDetail, error) {
	const query = `UPDATE users_details u SET fullname = $1, id_gender = $2, phone = $3, date_of_birth = $4, updated_at = $5 WHERE id_user = $6 AND deleted_at IS NULL RETURNING created_at, updated_at`

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

func (repo *UserDetailRepositoryImpl) FindByAccountId(c context.Context, id_account int64) (*pb.UserDetail, error) {
	query := `SELECT id, id_user, fullname, id_gender, phone, date_of_birth, created_at, updated_at, deleted_at FROM users_details WHERE id_user = $1 AND deleted_at IS NULL`

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

func (repo *UserDetailRepositoryImpl) FindAll(c context.Context) ([]*pb.UserDetail, error) {
	const query = `SELECT id, id_user, fullname, id_gender, phone, date_of_birth, created_at, updated_at, deleted_at FROM users_details WHERE deleted_at IS NULL`
	stmt, err := repo.DB.PrepareContext(c, query)
	if err != nil {
		return nil, err
	}

	users_details := make([]*pb.UserDetail, 0)
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

		users_details = append(users_details, &pb.UserDetail{
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

	return users_details, nil
}

func (repo *UserDetailRepositoryImpl) IsExist(c context.Context, id_account int64) (bool, error) {
	const query = `SELECT id FROM users_details u WHERE u.id_user = $1`
	var id int64

	log.Println(query)
	stmt, err := repo.DB.PrepareContext(c, query)
	if err != nil {
		log.Println("Error preparing context: ", err)
		return false, err
	}

	err = stmt.QueryRowContext(c, id_account).Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		log.Println("Error querying row: ", err)
		return false, err
	} else if err == sql.ErrNoRows {
		return false, nil
	} else {
		return true, nil
	}
}
