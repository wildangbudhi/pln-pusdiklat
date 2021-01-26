package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

func (handler *ForumHTTPHandler) GetForum(c *gin.Context) {

	c.Header("Content-Type", "application/json")

	forumIDString := c.Param("forum_id")

	forumID, err := domain.NewUUIDFromString(forumIDString)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	requestUserIDSFromHeader, ok := c.Request.Header["X-Auth-Id"]

	if !ok {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: "Request User ID Not Found"})
		return
	}

	requestUserID, err := strconv.Atoi(requestUserIDSFromHeader[0])

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: "Request User ID Format Invalid"})
		return
	}

	forumData, err := handler.forumUsecase.GetForum(*forumID, requestUserID)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, forumData)

}
