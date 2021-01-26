package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

type createForumRequestBody struct {
	Title      string `json:"title" binding:"required"`
	Question   string `json:"question" binding:"required"`
	CategoryID int    `json:"category_id" binding:"required"`
}

type createFroumResponseBody struct {
	ForumID string `json:"forum_id"`
}

func (handler *ForumHTTPHandler) CreateForum(c *gin.Context) {

	c.Header("Content-Type", "application/json")

	requestBodyData := &createForumRequestBody{}

	err := c.BindJSON(requestBodyData)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	requestUserIDSFromHeader, ok := c.Request.Header["X-Auth-Id"]

	if !ok {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: "Request User ID Not Found"})
		return
	}

	requestUserID, err := strconv.Atoi(requestUserIDSFromHeader[0])

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: "Request User ID Format Invalid"})
		return
	}

	forumID, err := handler.forumUsecase.CreateForum(requestBodyData.Title, requestBodyData.Question, requestUserID, requestBodyData.CategoryID)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, createFroumResponseBody{ForumID: forumID.GetValue()})

}
