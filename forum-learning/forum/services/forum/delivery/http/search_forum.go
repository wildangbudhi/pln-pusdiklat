package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

type searchForumRequestQuery struct {
	Offest *int   `form:"offset" json:"offset" binding:"required"`
	Limit  int    `form:"limit" json:"limit" binding:"required"`
	Query  string `form:"query" json:"query"`
}

type searchForumRequestHeader struct {
	XAuthID int `header:"X-Auth-Id" json:"X-Auth-Id" binding:"required"`
}

type searchForumResponseBody struct {
	Forum []domain.Forum `json:"forum" binding:"required"`
}

func (handler *ForumHTTPHandler) SearchForum(c *gin.Context) {

	c.Header("Content-Type", "application/json")

	var err error
	var forumData []domain.Forum

	requestHeader := &searchForumRequestHeader{}
	requestQuery := &searchForumRequestQuery{}

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

	forumData, err = handler.forumUsecase.SearchForum(*requestQuery.Offest, requestQuery.Limit, requestHeader.XAuthID, requestQuery.Query)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, searchForumResponseBody{Forum: forumData})

}
