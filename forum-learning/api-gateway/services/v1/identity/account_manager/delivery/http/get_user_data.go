package http

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/identity/account_manager/domain"
)

type getUserDataRequestHeader struct {
	Authorization string `header:"Authorization" json:"Authorization" binding:"required"`
}

func (handler *AccountManagerHTTPHandler) GetUserData(c *gin.Context) {

	requestHeader := &getUserDataRequestHeader{}

	err := c.BindHeader(requestHeader)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	userIDString := c.Param("id")

	userID, err := strconv.Atoi(userIDString)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	resp, err := handler.accountManagerUsecase.GetUserData(requestHeader.Authorization, userID, c.Request.RequestURI, c.Request.Method)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	for name, value := range resp.Header {
		c.Header(name, value[0])
	}

	contentType := resp.Header.Get("Content-Type")
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	if contentType == "" {
		contentType = "application/json"
	}

	c.Data(resp.StatusCode, contentType, body)

}
