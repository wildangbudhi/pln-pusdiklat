package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

type reactForumReactionRequestQuery struct {
	ReactionKey     string `form:"reaction_key" json:"reaction_key" binding:"required"`
	ReactionToggled *bool  `form:"reaction_toggled" json:"reaction_toggled" binding:"required"`
}

type reactForumReactionRequestHeader struct {
	XAuthID int `header:"X-Auth-Id" json:"X-Auth-Id" binding:"required"`
}

type reactForumReactionResponseBody struct {
	Status string `json:"status" binding:"required"`
}

func (handler *ForumHTTPHandler) ReactForumReplies(c *gin.Context) {

	c.Header("Content-Type", "application/json")

	var err error

	requestHeader := &reactForumReactionRequestHeader{}
	requestQuery := &reactForumReactionRequestQuery{}

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

	forumRepliesIDString := c.Param("forum_reply_id")

	forumRepliesID, err := domain.NewUUIDFromString(forumRepliesIDString)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	forumReactionType, err := domain.NewForumRepliesReactionType(requestQuery.ReactionKey, *requestQuery.ReactionToggled)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	err = handler.forumUsecase.ReactForumReplies(requestHeader.XAuthID, *forumRepliesID, *forumReactionType)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, reactForumReactionResponseBody{Status: "Reaction Successfully Recorded"})

}
