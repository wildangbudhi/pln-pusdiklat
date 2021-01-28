package usecase

import (
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

func (usecase *forumUsecase) FetchWithPaginationByAuthorID(requestUserID int, authorID int, offset int, limit int, topForumSort bool) ([]domain.Forum, error) {

	var forumData []domain.Forum
	var err error

	forumData, err = usecase.forumRepository.FetchForumByAuthorIDWithUserReaction(authorID, offset, limit, requestUserID, topForumSort)

	if err != nil {
		return nil, err
	}

	return forumData, nil

}
