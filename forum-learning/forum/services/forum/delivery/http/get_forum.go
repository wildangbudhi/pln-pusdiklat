package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

type getForumRequestHeader struct {
	XAuthID int `header:"X-Auth-Id" json:"X-Auth-Id" binding:"required"`
}

func (handler *ForumHTTPHandler) GetForum(c *gin.Context) {

	c.Header("Content-Type", "application/json")

	var err error
	var forumData *domain.Forum

	requestHeader := &fetchWithPaginationRequestHeader{}
	forumIDString := c.Param("forum_id")

	forumID, err := domain.NewUUIDFromString(forumIDString)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	err = c.BindHeader(requestHeader)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	forumData, err = handler.forumUsecase.GetForum(*forumID, requestHeader.XAuthID)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, forumData)

}
