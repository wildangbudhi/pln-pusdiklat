package usecase

import "github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"

func (usecase *forumUsecase) FetchReplyByForumIDWithPagination(requestUserID int, offset int, limit int, forumID domain.UUID) ([]domain.ForumReplies, error) {

	var forumRepliesData []domain.ForumReplies
	var err error

	forumRepliesData, err = usecase.forumRepliesRepository.FetchForumRepliesByForumIDWithUserReaction(offset, limit, forumID, requestUserID)

	if err != nil {
		return nil, err
	}

	return forumRepliesData, nil

}
