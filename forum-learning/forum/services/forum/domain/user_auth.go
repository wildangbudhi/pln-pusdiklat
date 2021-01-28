package domain

import (
	"database/sql"
	"encoding/json"
)

// UserAuth is Struct Model for user_auth table in database
type UserAuth struct {
	ID         int
	FullName   sql.NullString
	AvatarFile sql.NullString
	Email      Email
	Username   string
}

// MarshalJSON Custom Marshal Encoder for UserAuth Object
func (obj *UserAuth) MarshalJSON() ([]byte, error) {

	newObject := &struct {
		ID         int    `json:"id"`
		FullName   string `json:"full_name"`
		AvatarFile string `json:"avatar_file"`
		Email      string `json:"email"`
		Username   string `json:"usename"`
	}{
		ID:         obj.ID,
		FullName:   obj.FullName.String,
		AvatarFile: obj.AvatarFile.String,
		Email:      obj.Email.GetValue(),
		Username:   obj.Username,
	}

	return json.Marshal(newObject)

}

// UserAuthRepository Interface is a contract of Repository for User Auth Table
type UserAuthRepository interface {
	GetUserAuthByID(id int) (*UserAuth, error)
	UpdateUserAuthByID(id int, fullName string) (int, error)
	InsertUserAuth(userAuth *UserAuth) (int64, error)
	DeleteUserAuthByID(id int) (int, error)
}
