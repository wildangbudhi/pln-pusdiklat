package usecase

import (
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/event_consumer/domain"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/event_consumer/domain/model"
)

type eventConsumerUsecase struct {
	userAuthRepository model.UserAuthRepository
}

// NewEventConsumerUsecase is a Constructor of eventConsumerUsecase
// Which implement EventConsumerUsecase Interface
func NewEventConsumerUsecase(userAuthRepository model.UserAuthRepository) domain.EventConsumerUsecase {
	return &eventConsumerUsecase{
		userAuthRepository: userAuthRepository,
	}
}
