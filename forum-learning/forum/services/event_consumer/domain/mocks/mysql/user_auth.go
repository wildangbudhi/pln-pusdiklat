package mysql

import (
	"github.com/stretchr/testify/mock"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/event_consumer/domain/model"
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

func (repo *UserAuthRepositoryMock) InsertUserAuth(userAuth *model.UserAuth) (int64, error) {

	args := repo.Called()
	result := args.Get(0)
	return result.(int64), args.Error(1)

}

func (repo *UserAuthRepositoryMock) DeleteUserAuthByID(id int) (int, error) {

	args := repo.Called()
	result := args.Get(0)
	return result.(int), args.Error(1)

}
