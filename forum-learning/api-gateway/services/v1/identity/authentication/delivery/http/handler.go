package http

import (
	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/identity/authentication/domain"
)

type AuthenticationHTTPHandler struct {
	authenticationUsecase domain.AuthenticationUsecase
}

func NewAuthenticationHTTPHandler(router *gin.RouterGroup, authenticationUsecase domain.AuthenticationUsecase) {

	handler := AuthenticationHTTPHandler{
		authenticationUsecase: authenticationUsecase,
	}

	router.POST("/register", handler.Register)
	router.POST("/authenticate", handler.Authenticate)
	router.POST("/verify", handler.Verify)
	router.POST("/endpoint/authorize", handler.EndpointAuthorize)

}
