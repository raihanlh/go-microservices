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
		"VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?\\) RETURNING id, fullname, id_gender, phone, date_of_birth, created_at, updated_at"

	prep := mock.ExpectPrepare(query)
	prep.ExpectQuery().WithArgs(u.IdAccount, u.Fullname, u.IdGender, u.Phone, dob, u.CreatedAt.AsTime(), u.UpdatedAt.AsTime(), nil).WillReturnRows(rows)
	repo := repository.NewUserRepository(db)

	user_detail, err := repo.Save(context.TODO(), u)

	fmt.Println(user_detail)
	assert.NoError(t, err)
	assert.NotNil(t, user_detail)
}

func TestUpdate(t *testing.T) {
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

	query := "UPDATE user_details u SET fullname = \\?, id_gender = \\?, phone = \\?, date_of_birth = \\?, updated_at = \\? WHERE id_user = \\? AND deleted_at IS NULL RETURNING created_at, updated_at"

	prep := mock.ExpectPrepare(query)
	prep.ExpectQuery().WithArgs(u_updated.Fullname, u_updated.IdGender, u_updated.Phone, dob, u_updated.UpdatedAt.AsTime().In(loc), u_updated.IdAccount).WillReturnRows(rows)
	repo := repository.NewUserRepository(db)

	user_detail, err := repo.Update(context.TODO(), u_updated)
	fmt.Println(user_detail)
	assert.NoError(t, err)
	assert.NotNil(t, user_detail)
}

// func TestGetByIdt *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	loc := time.FixedZone("UTC+7", 7*60*60)

// 	rows := sqlmock.NewRows([]string{"id", "id_user", "fullname", "id_gender", "phone", "date_of_birth", "created_at", "updated_at", "deleted_at"}).
// 		AddRow(1, 1, "Test McTester", 0, "08123456789", time.Date(1990, time.January, 1, 0, 0, 0, 0, loc), time.Now(), time.Now(), nil)
// }
