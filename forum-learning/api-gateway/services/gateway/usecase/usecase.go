package usecase

import (
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/gateway/domain"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/gateway/domain/model"
)

type gatewayUsecase struct {
	enpointsRepository model.EnpointsRepository
}

func NewGatewayUsecase(enpointsRepository model.EnpointsRepository) domain.GatewayUsecase {
	return &gatewayUsecase{
		enpointsRepository: enpointsRepository,
	}
}
