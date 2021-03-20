package usecase

import "github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/identity/account_manager/domain"

func (usecase *accountManagerUsecase) Authenticate(token string) (*domain.VerifyResponse, error) {

	verifyRespones, err := usecase.authenticationRepository.Verify(token)

	if err != nil {
		return nil, err
	}

	return verifyRespones, nil

}
