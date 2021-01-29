package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

type replyForumRequestBody struct {
	Answer string `json:"answer" binding:"required"`
}

type replyForumRequestHeader struct {
	XAuthID int `header:"X-Auth-Id" json:"X-Auth-Id" binding:"required"`
}

type replyForumResponseBody struct {
	ForumRepiesID string `json:"forum_replies_id"`
}

func (handler *ForumHTTPHandler) ReplyForum(c *gin.Context) {

	c.Header("Content-Type", "application/json")

	requestHeader := &replyForumRequestHeader{}
	requestBodyData := &replyForumRequestBody{}

	forumIDString := c.Param("forum_id")

	forumID, err := domain.NewUUIDFromString(forumIDString)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	err = c.BindJSON(requestBodyData)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	err = c.BindHeader(requestHeader)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	forumRepliesID, err := handler.forumUsecase.ReplyForum(requestHeader.XAuthID, *forumID, requestBodyData.Answer)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, replyForumResponseBody{ForumRepiesID: forumRepliesID.GetValue()})

}
