package usecase

import "github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"

func (usecase *forumUsecase) FetchWithPagination(requestUserID int, offset int, limit int, categoriID *int, topForumSort bool) ([]domain.Forum, error) {

	var forumData []domain.Forum
	var err error

	if categoriID == nil {
		forumData, err = usecase.forumRepository.FetchForumWithUserReaction(offset, limit, requestUserID, topForumSort)
	} else {
		forumData, err = usecase.forumRepository.FetchForumByCategoryIDWithUserReaction(*categoriID, offset, limit, requestUserID, topForumSort)
	}

	if err != nil {
		return nil, err
	}

	return forumData, nil

}
