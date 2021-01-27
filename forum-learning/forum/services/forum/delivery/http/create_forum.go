package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

type createForumRequestBody struct {
	Title      string `json:"title" binding:"required"`
	Question   string `json:"question" binding:"required"`
	CategoryID int    `json:"category_id" binding:"required"`
}

type createFroumRequestHeader struct {
	XAuthID int `header:"X-Auth-Id" json:"X-Auth-Id" binding:"required"`
}

type createFroumResponseBody struct {
	ForumID string `json:"forum_id"`
}

func (handler *ForumHTTPHandler) CreateForum(c *gin.Context) {

	c.Header("Content-Type", "application/json")

	requestHeader := &createFroumRequestHeader{}
	requestBodyData := &createForumRequestBody{}

	err := c.BindJSON(requestBodyData)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	err = c.BindHeader(requestHeader)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	forumID, err := handler.forumUsecase.CreateForum(requestBodyData.Title, requestBodyData.Question, requestHeader.XAuthID, requestBodyData.CategoryID)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, createFroumResponseBody{ForumID: forumID.GetValue()})

}
