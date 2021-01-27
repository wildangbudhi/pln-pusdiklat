package usecase

import "github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"

func (usecase *forumUsecase) ReplyForum(requestUserID int, forumID domain.UUID, answer string) (*domain.UUID, error) {

	forumReplyID := domain.NewUUID()

	_, err := usecase.forumRepliesRepository.InsertForumReplies(*forumReplyID, requestUserID, forumID, answer)

	if err != nil {
		return nil, err
	}

	return forumReplyID, nil

}
