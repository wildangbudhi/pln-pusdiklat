package http

import (
	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

type ForumHTTPHandler struct {
	forumUsecase domain.ForumUsecase
}

func NewForumHTTPHandler(router *gin.RouterGroup, forumUsecase domain.ForumUsecase) {

	handler := ForumHTTPHandler{
		forumUsecase: forumUsecase,
	}

	router.GET("/category", handler.FetchCategory)
	router.POST("/create", handler.CreateForum)

}
