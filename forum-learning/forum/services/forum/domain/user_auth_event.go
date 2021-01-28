package domain

// UserAuthEventPayload is Struct Model for user_auth table in database
type UserAuthEventPayload struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

// UserAuthEvent is Struct for Event
type UserAuthEvent struct {
	Action string                `json:"action"` // CREATE, UPDATE, DELETE
	Data   *UserAuthEventPayload `json:"data,omitempty"`
}
