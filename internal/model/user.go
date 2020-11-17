package model

import "time"

type User struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	Password       string    `json:"password"`
	Occupation     string    `json:"occupation"`
	HashPassword   string    `json:"hash_password"`
	Email          string    `json:"email"`
	AvatarFileName string    `json:"avatar_file_name"`
	Role           string    `json:"role"`
	Token          string    `json:"token"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
