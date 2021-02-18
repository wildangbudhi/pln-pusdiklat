package usecase

import "github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/forum/domain"

func (usecase *forumUsecase) Authenticate(token string) (*domain.VerifyResponse, error) {

	verifyRespones, err := usecase.authenticationRepository.Verify(token)

	if err != nil {
		return nil, err
	}

	return verifyRespones, nil

}
