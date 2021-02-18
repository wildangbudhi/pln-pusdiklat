package usecase

import "github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/identity/account_manager/domain"

type accountManagerUsecase struct {
	accountManagerRepository domain.AccountManagerRepository
	authenticationRepository domain.AuthenticationRepository
}

// NewAccountManagerUsecase is a Constructor of accountManagerUsecase
// Which implement AccountManagerUsecase Interface
func NewAccountManagerUsecase(accountManagerRepository domain.AccountManagerRepository, authenticationRepository domain.AuthenticationRepository) domain.AccountManagerUsecase {
	return &accountManagerUsecase{
		accountManagerRepository: accountManagerRepository,
		authenticationRepository: authenticationRepository,
	}
}
