package usecase

import "github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"

func (usecase *forumUsecase) GetForum(forumID domain.UUID, requestUserID int) (*domain.Forum, error) {

	forum, err := usecase.forumRepository.GetForumByIDWithUserReaction(forumID, requestUserID)

	if err != nil {
		return nil, err
	}

	return forum, nil

}
