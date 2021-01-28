package usecase

import (
	"log"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

func (usecase *forumUsecase) SearchForum(offset int, limit int, requestUserID int, query string) ([]domain.Forum, error) {

	log.Printf("Start Usecase")

	forumData, err := usecase.forumRepository.SearchByTitleAndQuestionWithUserReaction(offset, limit, requestUserID, query)

	if err != nil {
		return nil, err
	}

	return forumData, nil

}
