package http

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

type deleteForumRequestHeader struct {
	XAuthID    int      `header:"X-Auth-Id" json:"X-Auth-Id" binding:"required"`
	XAuthRoles []string `header:"X-Auth-Roles" json:"X-Auth-Roles" binding:"required"`
}

type deleteForumResponseBody struct {
	Status string `json:"status" binding:"required"`
}

func (handler *ForumHTTPHandler) DeleteForum(c *gin.Context) {

	c.Header("Content-Type", "application/json")

	var err error

	requestHeader := &deleteForumRequestHeader{}

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

	forumIDString := c.Param("forum_id")

	forumID, err := domain.NewUUIDFromString(forumIDString)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	err = handler.forumUsecase.DeleteForum(*forumID, requestHeader.XAuthID, requestHeader.XAuthRoles)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, deleteForumResponseBody{Status: "Forum Successfully Deleted"})

}
