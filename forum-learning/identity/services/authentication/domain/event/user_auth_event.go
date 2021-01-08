package event

import "database/sql"

// Roles is Struct Model for roles table in database
type Roles struct {
	ID       int    `json:"id"`
	RoleName string `json:"role_name"`
}

// UserAuth is Struct Model for user_auth table in database
type UserAuth struct {
	ID       int            `json:"id"`
	FullName sql.NullString `json:"full_name"`
	Email    string         `json:"email"`
	Username string         `json:"username"`
	Roles    []Roles        `json:"roles"`
}

// UserAuthEvent is Struct for Event
type UserAuthEvent struct {
	Action string    `json:"action"` // CREATE, UPDATE, DELETE
	Data   *UserAuth `json:"data,omitempty"`
}

// UserAuthEventRepository is a contract of Repository for Message Queue UserAuthEvent
type UserAuthEventRepository interface {
	PublishDataChangesEvent(userAuthEvent *UserAuthEvent) error
}
