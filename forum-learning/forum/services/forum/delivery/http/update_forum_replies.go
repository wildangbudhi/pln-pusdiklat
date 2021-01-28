package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

type updateForumRepliesRequestHeader struct {
	XAuthID    int      `header:"X-Auth-Id" json:"X-Auth-Id" binding:"required"`
	XAuthRoles []string `header:"X-Auth-Roles" json:"X-Auth-Roles" binding:"required"`
}

type updateForumRepliesRequestBody struct {
	Answer string `json:"answer" binding:"required"`
}

type updateForumRepliesResponseBody struct {
	Status string `json:"status"`
}

func (handler *ForumHTTPHandler) UpdateForumReplies(c *gin.Context) {

	log.Printf("Start Handler")

	c.Header("Content-Type", "application/json")

	requestHeader := &updateForumRepliesRequestHeader{}
	requestBody := &updateForumRepliesRequestBody{}

	forumRepliesIDString := c.Param("forum_reply_id")

	forumRepliesID, err := domain.NewUUIDFromString(forumRepliesIDString)

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

	err = handler.forumUsecase.UpdateForumReplies(requestHeader.XAuthID, *forumRepliesID, requestBody.Answer, requestHeader.XAuthRoles)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, updateForumRepliesResponseBody{Status: "Forum Reply Successfully Updated"})

}
