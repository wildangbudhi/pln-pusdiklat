package event

// UserAuth is Struct Model for user_auth table in database
type UserAuth struct {
	ID         int    `json:"id"`
	FullName   string `json:"full_name"`
	AvatarFile string `json:"avatar_file"`
	Email      string `json:"email"`
	Username   string `json:"username"`
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
