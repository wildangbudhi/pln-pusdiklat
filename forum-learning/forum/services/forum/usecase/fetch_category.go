package usecase

import "github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"

func (usecase *forumUsecase) FetchCategory() ([]domain.Category, error) {

	categoryList, err := usecase.categoryRepository.FetchCategory()

	if err != nil {
		return nil, err
	}

	return categoryList, nil

}
