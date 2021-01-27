package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

type reactForumRequestQuery struct {
	ReactionKey     string `form:"reaction_key" json:"reaction_key" binding:"required"`
	ReactionToggled *bool  `form:"reaction_toggled" json:"reaction_toggled" binding:"required"`
}

type reactForumRequestHeader struct {
	XAuthID int `header:"X-Auth-Id" json:"X-Auth-Id" binding:"required"`
}

type reactForumResponseBody struct {
	Status string `json:"status" binding:"required"`
}

func (handler *ForumHTTPHandler) ReactForum(c *gin.Context) {

	c.Header("Content-Type", "application/json")

	var err error

	requestHeader := &reactForumRequestHeader{}
	requestQuery := &reactForumRequestQuery{}

	err = c.BindHeader(requestHeader)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	err = c.BindQuery(requestQuery)

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

	forumReactionType, err := domain.NewForumReactionType(requestQuery.ReactionKey, *requestQuery.ReactionToggled)

	err = handler.forumUsecase.ReactForum(requestHeader.XAuthID, *forumID, *forumReactionType)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, reactForumResponseBody{Status: "Reaction Successfully Recorded"})

}
