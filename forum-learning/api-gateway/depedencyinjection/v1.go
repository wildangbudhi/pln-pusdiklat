package depedencyinjection

import (
	"github.com/gin-gonic/gin"
	identityAuthenticationHTTPDelivery "github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/identity/authentication/delivery/http"
	identityAuthenticationHTTPRepository "github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/identity/authentication/repository/http"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/utils"
)

func identityAuthentticationDI(router *gin.RouterGroup, server *utils.Server) {

	authenticationRepository := identityAuthenticationHTTPRepository.NewAuthenticationRepository(server.HTTPClient, server.Config.EndpointsMap)
	identityAuthenticationHTTPDelivery.NewAuthenticationHTTPHandler(router, authenticationRepository)

}

func identityDI(router *gin.RouterGroup, server *utils.Server) {

	authenticationRouter := router.Group("auth")

	identityAuthentticationDI(authenticationRouter, server)

}

func ApiV1(server *utils.Server) {

	v1Route := server.APIRouter.Group("v1")

	identityRoute := v1Route.Group("identity")

	identityDI(identityRoute, server)

}
