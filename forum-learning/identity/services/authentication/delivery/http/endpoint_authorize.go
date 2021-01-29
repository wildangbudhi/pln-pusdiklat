package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain"
)

type endpointAuthorizeRequestBody struct {
	UserID int    `json:"user_id" binding:"required"`
	Method string `json:"method" binding:"required"`
	URL    string `json:"url" binding:"required"`
}

type endpointAuthorizeResponseBody struct {
	Authorized bool `json:"authorized" binding:"required"`
}

func (handler *AuthenticationHTTPHandler) EndpointAuthorize(c *gin.Context) {

	c.Header("Content-Type", "application/json")

	requestBodyData := &endpointAuthorizeRequestBody{}

	c.BindJSON(requestBodyData)

	if requestBodyData.Method == "" {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: "Method Cannot Be Empty"})
		return
	}

	if requestBodyData.URL == "" {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: "URL Cannot Be Empty"})
		return
	}

	authorizationStatus, err := handler.authenticationUsecase.EndpointAuthorize(
		requestBodyData.UserID,
		requestBodyData.Method,
		requestBodyData.URL,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, endpointAuthorizeResponseBody{Authorized: authorizationStatus})

}
