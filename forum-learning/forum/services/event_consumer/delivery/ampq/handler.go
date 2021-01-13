package ampq

import (
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/event_consumer/domain"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/utils"
)

type EventConsumerAMPQHandler struct {
	evenConsumerUsecase domain.EventConsumerUsecase
}

func NewEventConsumerAMPQHandler(server *utils.MessageQueueServer, evenConsumerUsecase domain.EventConsumerUsecase) {

	handler := EventConsumerAMPQHandler{
		evenConsumerUsecase: evenConsumerUsecase,
	}

	server.RegisterConsumerController("forum-queue", "UserAuthData", "event", "topic", handler.UserAuthDataChanges)

}
