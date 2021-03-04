package http

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/forum/domain"
)

func (handler *ForumHTTPHandler) ReactForumReplies(c *gin.Context) {

	requestURL := c.Request.RequestURI
	requestBody := c.Request.Body
	urlQuery := c.Request.URL.Query().Encode()
	requestHeader := c.Request.Header
	requestMethod := c.Request.Method
	forumRepliesID := c.Param("forum_reply_id")

	resp, err := handler.forumUsecase.ReactForumReplies(requestURL, requestMethod, forumRepliesID, requestBody, urlQuery, requestHeader)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	for name, values := range resp.Header {
		for i := 0; i < len(values); i++ {
			c.Header(name, values[i])
		}
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