package http

import (
	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/account_manager/domain"
)

type AccountManagerHTTPHandler struct {
	accountManagerUsecase domain.AccountManagerUsecase
}

func NewAccountManagerHTTPHandler(router *gin.RouterGroup, accountManagerUsecase domain.AccountManagerUsecase) {

	handler := AccountManagerHTTPHandler{
		accountManagerUsecase: accountManagerUsecase,
	}

	router.GET("/:id", handler.GetUserData)
	router.POST("/update/:id", handler.UpdateUserData)

}
