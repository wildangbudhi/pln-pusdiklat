package ampq

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/event_consumer/domain"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/event_consumer/domain/event"
)

func (handler *EventConsumerAMPQHandler) UserAuthDataChanges(msg amqp.Delivery) {

	msgObject := event.UserAuthEvent{}
	err := json.Unmarshal(msg.Body, &msgObject)

	if err != nil {
		log.Printf("[EventConsumer][UserAuthDataChanges] Error : %s\n", err.Error())
		return
	}

	email, err := domain.NewEmail(msgObject.Data.Email)

	if err != nil {
		log.Printf("[EventConsumer][UserAuthDataChanges] Error : %s\n", err.Error())
		return
	}

	err = handler.evenConsumerUsecase.UserAuthDataChangesEvent(msgObject.Action, msgObject.Data.ID, msgObject.Data.FullName, *email, msgObject.Data.Username)

	if err != nil {
		log.Printf("[EventConsumer][UserAuthDataChanges] Error : %s\n", err.Error())
		return
	}

	msg.Ack(false)

}
