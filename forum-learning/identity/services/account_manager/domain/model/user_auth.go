package model

import "database/sql"

// UserAuth is Struct Model for user_auth table in database
type UserAuth struct {
	ID         int            `json:"id"`
	FullName   sql.NullString `json:"full_name"`
	AvatarFile sql.NullString `json:"avatar_file"`
	Username   string         `json:"username"`
	EmployeeNo sql.NullString `json:"employee_no"`
	IsEmployee bool           `json:"is_employee"`
}

// UserAuthRepository Interface is a contract of Repository for User Auth Table
type UserAuthRepository interface {
	GetUserAuthByID(id int) (*UserAuth, error)
	UpdateUserAuthByID(id int, fullName string) (int, error)
}
