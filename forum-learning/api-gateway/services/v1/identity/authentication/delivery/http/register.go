package http

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/identity/authentication/domain"
)

func (handler *AuthenticationHTTPHandler) Register(c *gin.Context) {

	resp, err := handler.authenticationUsecase.Register(c.Request.Body)

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
