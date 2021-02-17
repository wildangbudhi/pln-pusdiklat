package http

import (
	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/identity/account_manager/domain"
)

type AccountManagerHTTPHandler struct {
	accountManagerUsecase domain.AccountManagerUsecase
}

func NewAuthenticationHTTPHandler(router *gin.RouterGroup, accountManagerUsecase domain.AccountManagerUsecase) {

	handler := AccountManagerHTTPHandler{
		accountManagerUsecase: accountManagerUsecase,
	}

	router.GET("/:id", handler.GetUserData)
	router.POST("/update/:id", handler.UpdateUserData)
}
