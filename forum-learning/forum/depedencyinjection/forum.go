package depedencyinjection

import (
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/delivery/ampq"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/delivery/http"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/repository/mysql"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/usecase"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/utils"
)

func ForumDI(server *utils.Server) {

	forumRoute := server.Router.Group("")

	userAuthRepository := mysql.NewUserAuthRepository(server.DB)
	categoryRepository := mysql.NewCategoryRepository(server.DB)
	forumRepository := mysql.NewForumRepository(server.DB)
	forumReactionRepository := mysql.NewForumReactionRepository(server.DB)
	forumRepliesRepository := mysql.NewForumRepliesRepository(server.DB)
	forumRepliesReactionRepository := mysql.NewForumRepliesReactionRepository(server.DB)

	usecase := usecase.NewForumUsecase(
		userAuthRepository,
		categoryRepository,
		forumRepository,
		forumReactionRepository,
		forumRepliesRepository,
		forumRepliesReactionRepository,
	)

	http.NewForumHTTPHandler(forumRoute, usecase)
	ampq.NewForumAMPQHandler(server.QueueServer, usecase)

}
