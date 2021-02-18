package http

import (
	"github.com/gin-gonic/gin"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/forum/domain"
)

type ForumHTTPHandler struct {
	forumUsecase domain.ForumUsecase
}

func NewAuthenticationHTTPHandler(router *gin.RouterGroup, forumUsecase domain.ForumUsecase) {

	handler := ForumHTTPHandler{
		forumUsecase: forumUsecase,
	}

	router.GET("/category", handler.FetchCategory)
}
