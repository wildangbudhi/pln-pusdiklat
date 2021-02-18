package usecase

import (
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain/model"
)

type authenticationUsecase struct {
	userAuthRepository model.UserAuthRepository
	secretKey          []byte
}

// NewAuthenticationUsecase is a Constructor of authenticationUsecase
// Which implement AuthenticationUsecase Interface
func NewAuthenticationUsecase(userAuthRepository model.UserAuthRepository, secretKey []byte) domain.AuthenticationUsecase {
	return &authenticationUsecase{
		userAuthRepository: userAuthRepository,
		secretKey:          secretKey,
	}
}
