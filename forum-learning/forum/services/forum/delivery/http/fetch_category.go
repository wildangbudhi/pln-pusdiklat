package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

type fetchCategoryResponseBody struct {
	Category []domain.Category `json:"category"`
}

func (handler *ForumHTTPHandler) FetchCategory(c *gin.Context) {

	c.Header("Content-Type", "application/json")

	categoryList, err := handler.forumUsecase.FetchCategory()

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.HTTPErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, fetchCategoryResponseBody{Category: categoryList})

}
