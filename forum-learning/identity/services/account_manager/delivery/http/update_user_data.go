package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/account_manager/domain"
)

type updateUserDataRequestBody struct {
	FullName string `json:"full_name" binding:"required"`
}

type updateUserDataResponseBody struct {
	Updated bool `json:"updated" binding:"required"`
}

func (handler *AccountManagerHTTPHandler) UpdateUserData(c *gin.Context) {

	c.Header("Content-Type", "application/json")

	userID := c.Param("id")
	userIDInt, err := strconv.Atoi(userID)

	requestBodyData := &updateUserDataRequestBody{}

	c.BindJSON(requestBodyData)

	requestUserID, ok := c.Request.Header["X-Auth-Id"]

	if !ok {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: "Request User ID Not Found"})
		return
	}

	requestUserIDInt, err := strconv.Atoi(requestUserID[0])

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: "Request User ID Format Invalid"})
		return
	}

	requestUserRoles, ok := c.Request.Header["X-Auth-Roles"]

	if !ok {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: "Request User Roles Not Found"})
		return
	}

	requestUserRolesArray := []string{}

	err = json.Unmarshal([]byte(requestUserRoles[0]), &requestUserRolesArray)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	status, err := handler.accountManagerUsecase.UpdateUserData(userIDInt, requestBodyData.FullName, requestUserIDInt, requestUserRolesArray)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusBadRequest, updateUserDataResponseBody{Updated: status})

}
