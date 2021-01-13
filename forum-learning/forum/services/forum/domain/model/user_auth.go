package model

import (
	"database/sql"
	"encoding/json"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

// UserAuth is Struct Model for user_auth table in database
type UserAuth struct {
	ID       int
	FullName sql.NullString
	Email    domain.Email
	Username string
}

// MarshalJSON Custom Marshal Encoder for UserAuth Object
func (obj *UserAuth) MarshalJSON() ([]byte, error) {

	newObject := &struct {
		ID       int    `json:"id"`
		FullName string `json:"full_name"`
		Email    string `json:"email"`
		Username string `json:"usename"`
	}{
		ID:       obj.ID,
		FullName: obj.FullName.String,
		Email:    obj.Email.GetValue(),
		Username: obj.Username,
	}

	return json.Marshal(newObject)

}
