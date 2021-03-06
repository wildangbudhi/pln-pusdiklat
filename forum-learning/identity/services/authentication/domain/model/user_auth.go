package model

import "database/sql"

// Roles is Struct Model for roles table in database
type Roles struct {
	ID       int    `json:"id"`
	RoleName string `json:"role_name"`
}

// UserAuth is Struct Model for user_auth table in database
type UserAuth struct {
	ID         int            `json:"id"`
	FullName   sql.NullString `json:"full_name"`
	Username   string         `json:"username"`
	Password   string         `json:"password"`
	Roles      []Roles        `json:"roles"`
	EmployeeNo sql.NullString `json:"employee_no"`
	IsEmployee bool           `json:"is_employee"`
}

// UserAuthRepository Interface is a contract of Repository for User Auth Table
type UserAuthRepository interface {
	GetUserAuthByID(id int) (*UserAuth, error)
	GetUserAuthByUsername(username string) (*UserAuth, error)
	InsertUserAuth(userAuth *UserAuth) (int64, error)
	CountRoleActivitiesPermission(method string, url string, roleIDList []int) (int, error)
}
