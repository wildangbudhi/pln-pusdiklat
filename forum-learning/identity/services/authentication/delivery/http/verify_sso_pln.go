package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain"
)

type verifySSOPLNRequestBody struct {
	Token string `json:"token" binding:"required"`
}

func (handler *AuthenticationHTTPHandler) VerifySSOPLN(c *gin.Context) {

	c.Header("Content-Type", "application/json")

	requestBodyData := &verifySSOPLNRequestBody{}

	c.BindJSON(requestBodyData)

	if requestBodyData.Token == "" {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: "Token Cannot Be Empty"})
		return
	}

	userAuthData, err := handler.authenticationUsecase.VerifySSOPLN(requestBodyData.Token)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, userAuthData)

}
