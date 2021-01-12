package usecase

import (
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/account_manager/domain"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/account_manager/domain/event"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/account_manager/domain/model"
)

type accountManagerUsecase struct {
	userAuthRepository      model.UserAuthRepository
	userAuthEventRepository event.UserAuthEventRepository
}

// NewAccountManagerUsecase is a Constructor of accountManagerUsecase
// Which implement AccountManagerUsecase Interface
func NewAccountManagerUsecase(userAuthRepository model.UserAuthRepository, userAuthEventRepository event.UserAuthEventRepository) domain.AccountManagerUsecase {
	return &accountManagerUsecase{
		userAuthRepository:      userAuthRepository,
		userAuthEventRepository: userAuthEventRepository,
	}
}
