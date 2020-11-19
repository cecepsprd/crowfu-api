package repository

import (
	"context"
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

func (u *userRepository) Get(ctx context.Context) ([]model.User, error) {
	rows, err := u.DB.QueryContext(ctx, "SELECT id, name, email, password, occupation, hash_password, avatar_file_name, role, token, created_at, updated_at FROM user")
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

func (u *userRepository) Save(c context.Context, user *model.User) (int64, error) {
	query := `INSERT INTO filter_word(name, email, password, occupation, hash_password, avatar_file_name, role, token, created_at, updated_at) VALUE (?,?,?,?,?,?,?,?,?,?)`

	res, err := u.DB.ExecContext(c, query)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	return res.RowsAffected()
}

func (u *userRepository) Update(c context.Context, id int64, user *model.User) (int64, error) {
	query := `UPDATE user SET name=?, email=?, password=?, occupation=?, hash_password=?, avatar_file_name=?, role=?, token=?, updated_at=? WHERE id = ?`

	res, err := u.DB.ExecContext(c, query)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	return res.RowsAffected()
}

func (u *userRepository) Delete(c context.Context, id int64) (int64, error) {
	query := `DELETE FROM user WHERE id = ?`

	res, err := u.DB.ExecContext(c, query)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	return res.RowsAffected()
}
