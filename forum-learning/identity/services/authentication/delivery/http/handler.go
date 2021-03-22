package http

import (
	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain"
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
	router.POST("/verify/sso/pln", handler.VerifySSOPLN)
	router.POST("/endpoint/authorize", handler.EndpointAuthorize)

}
