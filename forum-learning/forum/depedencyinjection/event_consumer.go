package depedencyinjection

import (
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/event_consumer/delivery/ampq"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/event_consumer/repository/mysql"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/event_consumer/usecase"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/utils"
)

func EventConsumerDI(server *utils.Server) {

	userAuthRepository := mysql.NewUserAuthRepository(server.DB)

	usecase := usecase.NewEventConsumerUsecase(userAuthRepository)

	ampq.NewEventConsumerAMPQHandler(server.QueueServer, usecase)

}
