package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

type fetchWithPaginationByAuthorIDRequestQuery struct {
	Offest       *int `form:"offset" json:"offset" binding:"required"`
	Limit        int  `form:"limit" json:"limit" binding:"required"`
	TopForumSort bool `form:"top_forum_sort" json:"top_forum_sort"`
}

type fetchWithPaginationByAuthorIDRequestHeader struct {
	XAuthID int `header:"X-Auth-Id" json:"X-Auth-Id" binding:"required"`
}

type fetchWithPaginationByAuthorIDResponseBody struct {
	Forum []domain.Forum `json:"forum" binding:"required"`
}

func (handler *ForumHTTPHandler) FetchWithPaginationByAuthorID(c *gin.Context) {

	c.Header("Content-Type", "application/json")

	var err error
	var forumData []domain.Forum

	requestHeader := &fetchWithPaginationByAuthorIDRequestHeader{}
	requestQuery := &fetchWithPaginationByAuthorIDRequestQuery{}

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

	authorIDString := c.Param("author_id")

	authorID, err := strconv.Atoi(authorIDString)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	forumData, err = handler.forumUsecase.FetchWithPaginationByAuthorID(requestHeader.XAuthID, authorID, *requestQuery.Offest, requestQuery.Limit, requestQuery.TopForumSort)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, fetchWithPaginationByAuthorIDResponseBody{Forum: forumData})

}
