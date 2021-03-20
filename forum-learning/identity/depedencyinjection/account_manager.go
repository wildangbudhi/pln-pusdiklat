package depedencyinjection

import (
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/account_manager/delivery/http"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/account_manager/repository/mysql"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/account_manager/usecase"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/utils"
)

func AccountManagerDI(server *utils.Server) {

	accountManagerRoute := server.Router.Group("user")

	userAuthRepository := mysql.NewUserAuthRepository(server.DB)

	usercase := usecase.NewAccountManagerUsecase(userAuthRepository)

	http.NewAccountManagerHTTPHandler(accountManagerRoute, usercase)

}
