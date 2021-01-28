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
	router.GET("/fetch", handler.FetchWithPagination)
	router.GET("/get/:forum_id", handler.GetForum)
	router.GET("/author/:author_id", handler.FetchWithPaginationByAuthorID)
	router.GET("/delete/:forum_id", handler.DeleteForum)
	router.GET("/search", handler.SearchForum)
	router.GET("/react/:forum_id", handler.ReactForum)
	router.POST("/reply/create/:forum_id", handler.ReplyForum)
	router.POST("/reply/update/:forum_reply_id", handler.UpdateForumReplies)

}
