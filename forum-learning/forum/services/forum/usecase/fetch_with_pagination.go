package usecase

import "github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"

func (usecase *forumUsecase) FetchWithPagination(requestUserID int, offset int, limit int, categoriID *int, topForumSort bool, timelineTimeFrame *domain.TimelineTimeFrame) ([]domain.Forum, error) {

	var forumData []domain.Forum
	var err error

	if categoriID == nil {
		forumData, err = usecase.forumRepository.FetchForumWithUserReaction(offset, limit, requestUserID, topForumSort, timelineTimeFrame)
	} else {
		forumData, err = usecase.forumRepository.FetchForumByCategoryIDWithUserReaction(*categoriID, offset, limit, requestUserID, topForumSort, timelineTimeFrame)
	}

	if err != nil {
		return nil, err
	}

	return forumData, nil

}
