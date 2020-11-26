package repository_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/cecepsprd/crowfu-api/internal/model"
	"github.com/cecepsprd/crowfu-api/internal/repository"
	"github.com/stretchr/testify/assert"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var NOW = time.Now()

var User = []model.User{
	model.User{
		ID:             10,
		Name:           "Elon Musk",
		Password:       "123",
		Occupation:     "Software Engineer",
		Email:          "elon@spacex.com",
		AvatarFileName: "elon.jpg",
		Role:           "superadmin",
		Token:          "token",
		CreatedAt:      NOW,
		UpdatedAt:      NOW,
	},
}

func TestGet(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println(err)
	}
	// qry := "SELECT id, name, email, password, occupation, hash_password, avatar_file_name, role, token, created_at, updated_at FROM user"
	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "occupation", "avatar_file_name", "role", "token", "created_at", "updated_at"}).
		AddRow(User[0].ID, User[0].Name, User[0].Email, User[0].Password, User[0].Occupation, User[0].AvatarFileName, User[0].Role, User[0].Token, NOW, NOW)

	mock.ExpectQuery("^SELECT (.+) FROM user").WillReturnRows(rows)

	repo := repository.NewUserRepository(db)
	result, err := repo.Get(context.TODO())

	assert.EqualValues(t, result, User)
	assert.NoError(t, err)
	// assert length of result
	assert.Len(t, result, 1)
}

func TestSave(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println(err)
	}

	qry := "INSERT INTO user\\(name, email, password, occupation, avatar_file_name, role, token\\) VALUE \\(\\?,\\?,\\?,\\?,\\?,\\?,\\?\\)"

	mock.ExpectExec(qry).
		WithArgs(User[0].Name, User[0].Email, User[0].Password, User[0].Occupation, User[0].AvatarFileName, User[0].Role, User[0].Token).
		WillReturnResult(sqlmock.NewResult(10, 1))

	repo := repository.NewUserRepository(db)
	rowsAffected, err := repo.Save(context.TODO(), &User[0])

	assert.NoError(t, err)
	assert.Equal(t, rowsAffected, int64(1))
}

func TestUpdate(t *testing.T) {
	var userID int64 = 10

	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println(err)
	}

	qry := "UPDATE user SET name=\\?, email=\\?, password=\\?, occupation=\\?, avatar_file_name=\\?, role=\\?, token=\\?, updated_at=\\? WHERE id = \\?"

	mock.ExpectExec(qry).
		WithArgs(User[0].Name, User[0].Email, User[0].Password, User[0].Occupation, User[0].AvatarFileName, User[0].Role, User[0].Token, User[0].UpdatedAt, userID).
		WillReturnResult(sqlmock.NewResult(userID, 1))

	repo := repository.NewUserRepository(db)
	rowsAffected, err := repo.Update(context.TODO(), userID, &User[0])

	assert.NoError(t, err)
	assert.Equal(t, rowsAffected, int64(1))
}

func TestDelete(t *testing.T) {
	var userID int64 = 10

	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println(err)
	}

	qry := "DELETE FROM user WHERE id = \\?"
	mock.ExpectExec(qry).
		WithArgs(userID).
		WillReturnResult(sqlmock.NewResult(userID, 1))

	repo := repository.NewUserRepository(db)
	rowsAffected, err := repo.Delete(context.TODO(), userID)

	assert.NoError(t, err)
	assert.Equal(t, rowsAffected, int64(1))
}
