package rabbitmq

import (
	"encoding/json"
	"fmt"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain/event"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/utils"
)

type userAuthEventRepository struct {
	queue *utils.MessageQueueServer
}

// NewUserAuthEventRepository is a constructor of userAuthEventRepository
// which implement UserAuthEventRepository Interface
func NewUserAuthEventRepository(queue *utils.MessageQueueServer) event.UserAuthEventRepository {
	return &userAuthEventRepository{
		queue: queue,
	}
}

func (repo *userAuthEventRepository) PublishDataChangesEvent(userAuthEvent *event.UserAuthEvent) error {

	var err error

	if userAuthEvent.Action == "" {
		return fmt.Errorf("Action of AuthUserEvent Cannot Be Empty")
	}

	jsonMessage, err := json.Marshal(userAuthEvent)

	if err != nil {
		return err
	}

	err = repo.queue.Publish("UserAuthData", "event", "topic", jsonMessage, "application/json")

	return err

}
