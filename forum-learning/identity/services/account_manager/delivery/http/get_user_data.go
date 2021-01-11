package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/account_manager/domain"
)

func (handler *AccountManagerHTTPHandler) GetUserData(c *gin.Context) {

	c.Header("Content-Type", "application/json")

	userID := c.Param("id")
	userIDInt, err := strconv.Atoi(userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: "User ID Format Invalid"})
		return
	}

	useAuthData, err := handler.accountManagerUsecase.GetUserData(userIDInt)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, useAuthData)

}
