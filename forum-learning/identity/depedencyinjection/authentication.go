package depedencyinjection

import (
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/delivery/http"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/repository/mysql"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/repository/rabbitmq"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/usecase"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/utils"
)

func AuthenticationDI(server *utils.Server) {

	authenticationRoute := server.Router.Group("auth")

	userAuthRepository := mysql.NewUserAuthRepository(server.DB)
	userAuthRepositoryEvent := rabbitmq.NewUserAuthEventRepository(server.QueueServer)

	usecase := usecase.NewAuthenticationUsecase(userAuthRepository, userAuthRepositoryEvent, server.Config.SecretKey)

	http.NewAuthenticationHTTPHandler(authenticationRoute, usecase)

}
