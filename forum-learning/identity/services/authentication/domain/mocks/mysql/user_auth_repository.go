package mysql

import (
	"github.com/stretchr/testify/mock"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain/model"
)

type UserAuthRepositoryMock struct {
	mock.Mock
}

func (repo *UserAuthRepositoryMock) GetUserAuthByID(id int) (*model.UserAuth, error) {

	args := repo.Called()
	result := args.Get(0)
	return result.(*model.UserAuth), args.Error(1)

}

func (repo *UserAuthRepositoryMock) GetUserAuthByEmail(email string) (*model.UserAuth, error) {

	args := repo.Called()
	result := args.Get(0)
	return result.(*model.UserAuth), args.Error(1)

}

func (repo *UserAuthRepositoryMock) GetUserAuthByUsername(username string) (*model.UserAuth, error) {

	args := repo.Called()
	result := args.Get(0)
	return result.(*model.UserAuth), args.Error(1)

}

func (repo *UserAuthRepositoryMock) InsertUserAuth(userAuth *model.UserAuth) (int64, error) {

	args := repo.Called()
	result := args.Get(0)
	return result.(int64), args.Error(1)

}

func (repo *UserAuthRepositoryMock) CountRoleActivitiesPermission(method string, url string, roleIDList []int) (int, error) {

	args := repo.Called()
	result := args.Get(0)
	return result.(int), args.Error(1)

}
