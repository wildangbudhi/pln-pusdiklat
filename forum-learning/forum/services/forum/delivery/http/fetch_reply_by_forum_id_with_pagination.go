package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

type fetchReplyByForumIDWithPaginationByAuthorIDRequestQuery struct {
	Offest       *int `form:"offset" json:"offset" binding:"required"`
	Limit        int  `form:"limit" json:"limit" binding:"required"`
	TopForumSort bool `form:"top_forum_sort" json:"top_forum_sort"`
}

type fetchReplyByForumIDWithPaginationByAuthorIDRequestHeader struct {
	XAuthID int `header:"X-Auth-Id" json:"X-Auth-Id" binding:"required"`
}

type fetchReplyByForumIDWithPaginationByAuthorIDResponseBody struct {
	ForumReplies []domain.ForumReplies `json:"forum_replies" binding:"required"`
}

func (handler *ForumHTTPHandler) FetchReplyByForumIDWithPagination(c *gin.Context) {

	c.Header("Content-Type", "application/json")

	var err error
	var forumRepliesData []domain.ForumReplies

	requestHeader := &fetchReplyByForumIDWithPaginationByAuthorIDRequestHeader{}
	requestQuery := &fetchReplyByForumIDWithPaginationByAuthorIDRequestQuery{}

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

	forumRepliesData, err = handler.forumUsecase.FetchReplyByForumIDWithPagination(requestHeader.XAuthID, *requestQuery.Offest, requestQuery.Limit, *forumID)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, fetchReplyByForumIDWithPaginationByAuthorIDResponseBody{ForumReplies: forumRepliesData})

}
