package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

type fetchWithPaginationRequestQuery struct {
	Offest       *int   `form:"offset" json:"offset" binding:"required"`
	Limit        int    `form:"limit" json:"limit" binding:"required"`
	CategoryID   *int   `form:"category_id" json:"category_id"`
	TopForumSort bool   `form:"top_forum_sort" json:"top_forum_sort"`
	TimeFrame    string `form:"time_frame" json:"time_frame"`
}

type fetchWithPaginationRequestHeader struct {
	XAuthID int `header:"X-Auth-Id" json:"X-Auth-Id" binding:"required"`
}

type fetchWithPaginationResponseBody struct {
	Forum []domain.Forum `json:"forum" binding:"required"`
}

func (handler *ForumHTTPHandler) FetchWithPagination(c *gin.Context) {

	c.Header("Content-Type", "application/json")

	var err error
	var forumData []domain.Forum

	requestHeader := &fetchWithPaginationRequestHeader{}
	requestQuery := &fetchWithPaginationRequestQuery{}

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

	var timelineTimeFrame *domain.TimelineTimeFrame = nil

	if requestQuery.TimeFrame != "" {

		timelineTimeFrame, err = domain.NewTimelineTimeFrame(requestQuery.TimeFrame)

		if err != nil {
			c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
			return
		}

	}

	forumData, err = handler.forumUsecase.FetchWithPagination(requestHeader.XAuthID, *requestQuery.Offest, requestQuery.Limit, requestQuery.CategoryID, requestQuery.TopForumSort, timelineTimeFrame)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, fetchWithPaginationResponseBody{Forum: forumData})

}
