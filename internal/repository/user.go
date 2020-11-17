package repository

import (
	"database/sql"

	"github.com/cecepsprd/crowfu-api/internal/model"
	"github.com/cecepsprd/crowfu-api/pkg/log"
)

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (u *userRepository) Get() ([]model.User, error) {
	rows, err := u.DB.Query("SELECT id, name, email, password, occupation, hash_password, avatar_file_name, role, token, created_at, updated_at FROM user")
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()

	users := make([]model.User, 0)

	for rows.Next() {
		user := model.User{}
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Occupation, &user.HashPassword, &user.AvatarFileName, &user.Role, &user.Token, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
