package usecase

import "github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/identity/authentication/domain"

type authenticationUsecase struct {
	authenticationRepository domain.AuthenticationRepository
}

// NewAuthenticationUsecase is a Constructor of authenticationUsecase
// Which implement AuthenticationUsecase Interface
func NewAuthenticationUsecase(authenticationRepository domain.AuthenticationRepository) domain.AuthenticationUsecase {
	return &authenticationUsecase{
		authenticationRepository: authenticationRepository,
	}
}
