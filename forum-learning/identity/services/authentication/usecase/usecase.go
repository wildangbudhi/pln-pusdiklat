package usecase

import (
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain/event"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain/model"
)

type authenticationUsecase struct {
	userAuthRepository      model.UserAuthRepository
	userAuthEventRepository event.UserAuthEventRepository
	secretKey               []byte
}

// NewAuthenticationUsecase is a Constructor of authenticationUsecase
// Which implement AuthenticationUsecase Interface
func NewAuthenticationUsecase(userAuthRepository model.UserAuthRepository, userAuthEventRepository event.UserAuthEventRepository, secretKey []byte) domain.AuthenticationUsecase {
	return &authenticationUsecase{
		userAuthRepository:      userAuthRepository,
		userAuthEventRepository: userAuthEventRepository,
		secretKey:               secretKey,
	}
}
