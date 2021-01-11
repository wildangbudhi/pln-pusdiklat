package event

import "github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/account_manager/domain/model"

// UserAuthEvent is Struct for Event
type UserAuthEvent struct {
	Action string          `json:"action"` // CREATE, UPDATE, DELETE
	Data   *model.UserAuth `json:"data,omitempty"`
}

// UserAuthEventRepository is a contract of Repository for Message Queue UserAuthEvent
type UserAuthEventRepository interface {
	PublishDataChangesEvent(userAuthEvent *UserAuthEvent) error
}
