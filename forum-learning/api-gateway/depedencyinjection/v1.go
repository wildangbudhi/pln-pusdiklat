package depedencyinjection

import (
	"github.com/gin-gonic/gin"
	forumHTTPDelivery "github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/forum/delivery/http"
	forumHTTPRepository "github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/forum/repository/http"
	forumUsecase "github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/forum/usecase"
	identityAccountManagerHTTPDelivery "github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/identity/account_manager/delivery/http"
	identityAccountManagerHTTPRepository "github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/identity/account_manager/repository/http"
	identityAccountManaerUsecase "github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/identity/account_manager/usecase"
	identityAuthenticationHTTPDelivery "github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/identity/authentication/delivery/http"
	identityAuthenticationHTTPRepository "github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/identity/authentication/repository/http"
	identityAuthenticationUsecase "github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/identity/authentication/usecase"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/utils"
)

func identityAuthenticationDI(router *gin.RouterGroup, server *utils.Server) {

	authenticationRepository := identityAuthenticationHTTPRepository.NewAuthenticationRepository(server.HTTPClient, server.Config.EndpointsMap)
	authenticationUsecase := identityAuthenticationUsecase.NewAuthenticationUsecase(authenticationRepository)
	identityAuthenticationHTTPDelivery.NewAuthenticationHTTPHandler(router, authenticationUsecase)

}

func identityAccountManagerDI(router *gin.RouterGroup, server *utils.Server) {

	authenticationRepository := identityAccountManagerHTTPRepository.NewAuthenticationRepository(server.HTTPClient, server.Config.EndpointsMap)
	accountManagerReposutory := identityAccountManagerHTTPRepository.NewAccountManagerRepository(server.HTTPClient, server.Config.EndpointsMap)
	accountManagerUsecase := identityAccountManaerUsecase.NewAccountManagerUsecase(accountManagerReposutory, authenticationRepository)
	identityAccountManagerHTTPDelivery.NewAuthenticationHTTPHandler(router, accountManagerUsecase)

}

func identityDI(router *gin.RouterGroup, server *utils.Server) {

	authenticationRouter := router.Group("auth")
	accountManagerRouter := router.Group("user")

	identityAuthenticationDI(authenticationRouter, server)
	identityAccountManagerDI(accountManagerRouter, server)

}

func forumDI(router *gin.RouterGroup, server *utils.Server) {

	authenticationRepository := forumHTTPRepository.NewAuthenticationRepository(server.HTTPClient, server.Config.EndpointsMap)
	forumReposutory := forumHTTPRepository.NewForumRepository(server.HTTPClient, server.Config.EndpointsMap)

	forumUsecase := forumUsecase.NewForumUsecase(forumReposutory, authenticationRepository)

	forumHTTPDelivery.NewAuthenticationHTTPHandler(router, forumUsecase)

}

func ApiV1(server *utils.Server) {

	v1Route := server.APIRouter.Group("v1")

	identityRoute := v1Route.Group("identity")
	forumRoute := v1Route.Group("forum")

	identityDI(identityRoute, server)
	forumDI(forumRoute, server)

}
