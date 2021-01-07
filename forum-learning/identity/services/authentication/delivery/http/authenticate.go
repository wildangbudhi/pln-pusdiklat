package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain"
)

type authenticateRequestBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type authenticateResponseBody struct {
	Token string `json:"token" binding:"required"`
}

// Authenticate is a HTTP Handler for Authenticate Usecase
func (handler *AuthenticationHTTPHandler) Authenticate(c *gin.Context) {

	c.Header("Content-Type", "application/json")

	requestBodyData := &authenticateRequestBody{}

	c.BindJSON(requestBodyData)

	if requestBodyData.Username == "" {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: "Username tidak boleh kosong"})
		return
	}

	if requestBodyData.Password == "" {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: "Password tidak boleh kosong"})
		return
	}

	token, err := handler.authenticationUsecase.Authenticate(requestBodyData.Username, requestBodyData.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, authenticateResponseBody{Token: token})
}
