package repository_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	pb "github.com/raihanlh/go-user-microservice/proto"
	"github.com/raihanlh/go-user-microservice/src/repository"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var current_time = time.Now()

var u = &pb.UserDetail{
	IdAccount: 1,
	Fullname:  "Test McTester",
	IdGender:  0,
	Gender:    "male",
	Phone:     "08123456789",
	DateOfBirth: &pb.Date{
		Year:  1990,
		Month: 1,
		Day:   1,
	},
	CreatedAt: timestamppb.New(current_time),
	UpdatedAt: timestamppb.New(current_time),
}

var u_updated = &pb.UserDetail{
	IdAccount: 1,
	Fullname:  "Test McTester Updated",
	IdGender:  1,
	Gender:    "female",
	Phone:     "08123456781",
	DateOfBirth: &pb.Date{
		Year:  1991,
		Month: 2,
		Day:   2,
	},
	CreatedAt: timestamppb.New(current_time),
	UpdatedAt: timestamppb.New(current_time),
}

func TestSave(t *testing.T) {
	fmt.Println("Test save")
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	loc := time.FixedZone("UTC+7", 7*60*60)

	dob := time.Date(int(u.DateOfBirth.Year), time.Month(int(u.DateOfBirth.Month)), int(u.DateOfBirth.Day), 0, 0, 0, 0, loc)

	rows := sqlmock.NewRows([]string{"id", "fullname", "id_gender", "phone", "date_of_birth", "created_at", "updated_at"}).
		AddRow(1, u.Fullname, u.IdGender, u.Phone, dob, u.CreatedAt.AsTime(), u.UpdatedAt.AsTime())

	query := "INSERT INTO user_details \\(id_user, fullname, id_gender, phone, date_of_birth, created_at, updated_at, deleted_at\\) " +
		"VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5, \\$6, \\$7, \\$8\\) RETURNING id, fullname, id_gender, phone, date_of_birth, created_at, updated_at"

	prep := mock.ExpectPrepare(query)
	prep.ExpectQuery().WithArgs(u.IdAccount, u.Fullname, u.IdGender, u.Phone, dob, u.CreatedAt.AsTime(), u.UpdatedAt.AsTime(), nil).WillReturnRows(rows)
	repo := repository.NewUserDetailRepository(db)

	user_detail, err := repo.Save(context.TODO(), u)

	fmt.Println(user_detail)
	assert.NoError(t, err)
	assert.NotNil(t, user_detail)
}

func TestUpdate(t *testing.T) {
	fmt.Println("Test update")
	updateTime := time.Now()
	u_updated.UpdatedAt = timestamppb.New(updateTime)

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	loc := time.FixedZone("UTC+7", 7*60*60)

	dob := time.Date(int(u_updated.DateOfBirth.Year), time.Month(int(u_updated.DateOfBirth.Month)), int(u_updated.DateOfBirth.Day), 0, 0, 0, 0, loc)

	rows := sqlmock.NewRows([]string{"created_at", "updated_at"}).
		AddRow(u_updated.CreatedAt.AsTime().In(loc), u_updated.UpdatedAt.AsTime().In(loc))

	query := "UPDATE user_details u SET fullname = \\$1, id_gender = \\$2, phone = \\$3, date_of_birth = \\$4, updated_at = \\$5 WHERE id_user = \\$6 AND deleted_at IS NULL RETURNING created_at, updated_at"

	prep := mock.ExpectPrepare(query)
	prep.ExpectQuery().WithArgs(u_updated.Fullname, u_updated.IdGender, u_updated.Phone, dob, u_updated.UpdatedAt.AsTime().In(loc), u_updated.IdAccount).WillReturnRows(rows)
	repo := repository.NewUserDetailRepository(db)

	user_detail, err := repo.Update(context.TODO(), u_updated)
	fmt.Println(user_detail)
	assert.NoError(t, err)
	assert.NotNil(t, user_detail)
}

func TestGetById(t *testing.T) {
	fmt.Println("Test get by id")
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	loc := time.FixedZone("UTC+7", 7*60*60)

	query := "SELECT id, id_user, fullname, id_gender, phone, date_of_birth, created_at, updated_at, deleted_at FROM user_details WHERE id_user = \\$1 AND deleted_at IS NULL"

	rows := sqlmock.NewRows([]string{"id", "id_user", "fullname", "id_gender", "phone", "date_of_birth", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, 1, "Test McTester", 0, "08123456789", time.Date(1990, time.January, 1, 0, 0, 0, 0, loc), time.Now(), time.Now(), nil)

	prep := mock.ExpectPrepare(query)
	prep.ExpectQuery().WithArgs(u.IdAccount).WillReturnRows(rows)
	repo := repository.NewUserDetailRepository(db)
	user_detail, err := repo.FindByAccountId(context.TODO(), u.IdAccount)
	fmt.Println(user_detail)
	assert.NoError(t, err)
	assert.NotNil(t, user_detail)
}

func TestGetAll(t *testing.T) {
	fmt.Println("Test get all")
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	loc := time.FixedZone("UTC+7", 7*60*60)

	query := "SELECT id, id_user, fullname, id_gender, phone, date_of_birth, created_at, updated_at, deleted_at FROM user_details WHERE deleted_at IS NULL"

	rows := sqlmock.NewRows([]string{"id", "id_user", "fullname", "id_gender", "phone", "date_of_birth", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, 1, "Test McTester", 0, "08123456789", time.Date(1990, time.January, 1, 0, 0, 0, 0, loc), time.Now(), time.Now(), nil).
		AddRow(2, 2, "Test McTester 2", 1, "08123456788", time.Date(1991, time.January, 1, 0, 0, 0, 0, loc), time.Now(), time.Now(), nil)

	prep := mock.ExpectPrepare(query)
	prep.ExpectQuery().WithArgs().WillReturnRows(rows)
	repo := repository.NewUserDetailRepository(db)
	user_details, err := repo.FindAll(context.TODO())
	fmt.Println(user_details)
	assert.NoError(t, err)
	assert.NotNil(t, user_details)
}

func TestIsExist(t *testing.T) {
	fmt.Println("Test is exist")
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	const query = "SELECT id FROM users_details u WHERE u.id_user = \\$1"

	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(1)

	prep := mock.ExpectPrepare(query)
	prep.ExpectQuery().WithArgs(u.IdAccount).WillReturnRows(rows)
	repo := repository.NewUserDetailRepository(db)
	exist, err := repo.IsExist(context.TODO(), 1)
	fmt.Println(exist)
	assert.NoError(t, err)
	assert.NotNil(t, exist)
}
