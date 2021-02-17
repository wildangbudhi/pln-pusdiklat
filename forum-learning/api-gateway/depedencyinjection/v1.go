package depedencyinjection

import (
	"github.com/gin-gonic/gin"
	identityHTTPDelivery "github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/identity/authentication/delivery/http"
	identityHTTPRepository "github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/identity/authentication/repository/http"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/utils"
)

func identityDI(router *gin.RouterGroup, server *utils.Server) {

	authenticationRepository := identityHTTPRepository.NewAuthenticationRepository(server.HTTPClient, server.Config.EndpointsMap)
	identityHTTPDelivery.NewAuthenticationHTTPHandler(router, authenticationRepository)

}

func ApiV1(server *utils.Server) {

	v1Route := server.APIRouter.Group("v1")

	identityRoute := v1Route.Group("identity")

	identityDI(identityRoute, server)

}
