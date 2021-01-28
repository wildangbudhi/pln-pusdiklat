package usecase

import (
	"fmt"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

func (usecase *forumUsecase) ReplyForum(requestUserID int, forumID domain.UUID, answer string) (*domain.UUID, error) {

	forumData, err := usecase.forumRepository.GetForumByIDWithUserReaction(forumID, requestUserID)

	if err != nil {
		return nil, err
	}

	if forumData.Status.String == "CLOSED" {
		return nil, fmt.Errorf("Forum Has Been Closed, You Cannot Reply This Forum Anymore")
	}

	forumReplyID := domain.NewUUID()

	_, err = usecase.forumRepliesRepository.InsertForumReplies(*forumReplyID, requestUserID, forumID, answer)

	if err != nil {
		return nil, err
	}

	return forumReplyID, nil

}
