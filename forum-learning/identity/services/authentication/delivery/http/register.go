package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain"
)

type registerRequestBody struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type registerResponseBody struct {
	UserAuthID int64 `json:"user_auth_id" binding:"required"`
}

// Register is a HTTP Handler for Register Usecase
func (handler *AuthenticationHTTPHandler) Register(c *gin.Context) {

	c.Header("Content-Type", "application/json")

	requestBodyData := &registerRequestBody{}

	c.BindJSON(requestBodyData)

	if requestBodyData.Email == "" {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: "Email Cannot Be Empty"})
		return
	}

	if requestBodyData.Username == "" {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: "Username Cannot Be Empty"})
		return
	}

	if requestBodyData.Password == "" {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: "Password Cannot Be Empty"})
		return
	}

	requestBodyData.Email = strings.ToLower(requestBodyData.Email)

	email, err := domain.NewEmail(requestBodyData.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	userAuthID, err := handler.authenticationUsecase.Register(requestBodyData.FullName, *email, requestBodyData.Username, requestBodyData.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, registerResponseBody{UserAuthID: userAuthID})

}
