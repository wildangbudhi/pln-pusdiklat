package http

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

type deleteForumRepliesRequestHeader struct {
	XAuthID    int      `header:"X-Auth-Id" json:"X-Auth-Id" binding:"required"`
	XAuthRoles []string `header:"X-Auth-Roles" json:"X-Auth-Roles" binding:"required"`
}

type deleteForumRepliesResponseBody struct {
	Status string `json:"status" binding:"required"`
}

func (handler *ForumHTTPHandler) DeleteForumReplies(c *gin.Context) {

	c.Header("Content-Type", "application/json")

	var err error

	requestHeader := &deleteForumRepliesRequestHeader{}

	err = c.BindHeader(requestHeader)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	err = json.Unmarshal([]byte(requestHeader.XAuthRoles[0]), &requestHeader.XAuthRoles)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	forumReplyIDString := c.Param("forum_reply_id")

	forumReplyID, err := domain.NewUUIDFromString(forumReplyIDString)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	err = handler.forumUsecase.DeleteForumReplies(requestHeader.XAuthID, requestHeader.XAuthRoles, *forumReplyID)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, deleteForumRepliesResponseBody{Status: "Forum Reply Successfully Deleted"})

}
