package mysql

import (
	"github.com/stretchr/testify/mock"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/account_manager/domain/model"
)

type UserAuthRepositoryMock struct {
	mock.Mock
}

func (repo *UserAuthRepositoryMock) GetUserAuthByID(id int) (*model.UserAuth, error) {

	args := repo.Called()
	result := args.Get(0)
	return result.(*model.UserAuth), args.Error(1)

}

func (repo *UserAuthRepositoryMock) UpdateUserAuthByID(id int, fullName string) (int, error) {

	args := repo.Called()
	result := args.Get(0)
	return result.(int), args.Error(1)

}
