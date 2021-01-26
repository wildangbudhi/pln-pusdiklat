package ampq

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

func (handler *ForumAMPQHandler) UserAuthDataChanges(msg amqp.Delivery) {

	msgObject := domain.UserAuthEvent{}
	err := json.Unmarshal(msg.Body, &msgObject)

	if err != nil {
		log.Printf("[Forum][UserAuthDataChanges] Error : %s\n", err.Error())
		return
	}

	email, err := domain.NewEmail(msgObject.Data.Email)

	if err != nil {
		log.Printf("[Forum][UserAuthDataChanges] Error : %s\n", err.Error())
		return
	}

	err = handler.forumUsecase.UserAuthDataChangesEvent(msgObject.Action, msgObject.Data.ID, msgObject.Data.FullName, *email, msgObject.Data.Username)

	if err != nil {
		log.Printf("[Forum][UserAuthDataChanges] Error : %s\n", err.Error())
		return
	}

	msg.Ack(false)

}
