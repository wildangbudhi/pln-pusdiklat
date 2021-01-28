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
	router.POST("/update/:forum_id", handler.UpdateForum)
	router.GET("/close/:forum_id", handler.CloseForum)
	router.GET("/delete/:forum_id", handler.DeleteForum)
	router.GET("/fetch", handler.FetchWithPagination)
	router.GET("/get/:forum_id", handler.GetForum)
	router.GET("/author/:author_id", handler.FetchWithPaginationByAuthorID)
	router.GET("/search", handler.SearchForum)
	router.GET("/react/:forum_id", handler.ReactForum)
	router.POST("/reply/create/:forum_id", handler.ReplyForum)
	router.POST("/reply/update/:forum_reply_id", handler.UpdateForumReplies)
	router.GET("/reply/delete/:forum_reply_id", handler.DeleteForumReplies)
	router.GET("/reply/react/:forum_reply_id", handler.ReactForumReplies)
	router.GET("/get/:forum_id/replies", handler.FetchReplyByForumIDWithPagination)

}
