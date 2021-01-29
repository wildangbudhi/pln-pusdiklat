package http

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

type updateForumRequestHeader struct {
	XAuthID    int      `header:"X-Auth-Id" json:"X-Auth-Id" binding:"required"`
	XAuthRoles []string `header:"X-Auth-Roles" json:"X-Auth-Roles" binding:"required"`
}

type updateForumRequestBody struct {
	Title      string `json:"title" binding:"required"`
	Question   string `json:"question" binding:"required"`
	CategoryID int    `json:"category_id" binding:"required"`
}

type updateForumResponseBody struct {
	Status string `json:"status"`
}

func (handler *ForumHTTPHandler) UpdateForum(c *gin.Context) {

	c.Header("Content-Type", "application/json")

	requestHeader := &updateForumRequestHeader{}
	requestBody := &updateForumRequestBody{}

	forumIDString := c.Param("forum_id")

	forumID, err := domain.NewUUIDFromString(forumIDString)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	err = c.BindJSON(requestBody)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

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

	err = handler.forumUsecase.UpdateForum(requestHeader.XAuthID, *forumID, requestBody.Title, requestBody.Question, requestBody.CategoryID, requestHeader.XAuthRoles)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, updateForumResponseBody{Status: "Forum Successfully Updated"})

}
