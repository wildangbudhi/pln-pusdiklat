package model

import (
	"database/sql"
)

// UserAuth is Struct Model for user_auth table in database
type UserAuth struct {
	ID         int            `json:"id"`
	FullName   sql.NullString `json:"full_name"`
	AvatarFile sql.NullString `json:"avatar_file"`
	Email      string         `json:"email"`
	Username   string         `json:"username"`
}

// UserAuthRepository Interface is a contract of Repository for User Auth Table
type UserAuthRepository interface {
	GetUserAuthByID(id int) (*UserAuth, error)
	UpdateUserAuthByID(id int, fullName string) (int, error)
	InsertUserAuth(userAuth *UserAuth) (int64, error)
	DeleteUserAuthByID(id int) (int, error)
}
