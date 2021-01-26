package mysql

import (
	"github.com/stretchr/testify/mock"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

type CategoryRepositoryMock struct {
	mock.Mock
}

func (repo *CategoryRepositoryMock) FetchCategory() ([]domain.Category, error) {

	args := repo.Called()
	result := args.Get(0)
	return result.([]domain.Category), args.Error(1)

}
