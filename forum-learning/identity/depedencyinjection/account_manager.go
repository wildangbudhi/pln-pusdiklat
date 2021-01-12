package depedencyinjection

import (
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/account_manager/delivery/http"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/account_manager/repository/mysql"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/account_manager/repository/rabbitmq"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/account_manager/usecase"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/utils"
)

func AccountManagerDI(server *utils.Server) {

	accountManagerRoute := server.Router.Group("user")

	userAuthRepository := mysql.NewUserAuthRepository(server.DB)
	userAuthRepositoryEvent := rabbitmq.NewUserAuthEventRepository(server.QueueServer)

	usercase := usecase.NewAccountManagerUsecase(userAuthRepository, userAuthRepositoryEvent)

	http.NewAccountManagerHTTPHandler(accountManagerRoute, usercase)

}
