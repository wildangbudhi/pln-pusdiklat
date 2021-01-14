package mysql

import (
	"github.com/stretchr/testify/mock"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain/model"
)

type CategoryRepositoryMock struct {
	mock.Mock
}

func (repo *CategoryRepositoryMock) FetchCategory() ([]model.Category, error) {

	args := repo.Called()
	result := args.Get(0)
	return result.([]model.Category), args.Error(1)

}
