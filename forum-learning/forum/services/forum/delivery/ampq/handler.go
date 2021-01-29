package ampq

import (
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/utils"
)

type ForumAMPQHandler struct {
	forumUsecase domain.ForumUsecase
}

func NewForumAMPQHandler(server *utils.MessageQueueServer, forumUsecase domain.ForumUsecase) {

	handler := ForumAMPQHandler{
		forumUsecase: forumUsecase,
	}

	server.RegisterConsumerController("forum-queue", "UserAuthData", "event", "topic", handler.UserAuthDataChanges)

}
