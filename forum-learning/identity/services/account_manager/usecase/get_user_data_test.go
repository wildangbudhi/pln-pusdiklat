package usecase_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/account_manager/domain/mocks/mysql"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/account_manager/domain/model"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/account_manager/usecase"
)

func TestGetUserDataSuccess(t *testing.T) {

	mockUserAuthRepository := new(mysql.UserAuthRepositoryMock)

	testService := usecase.NewAccountManagerUsecase(mockUserAuthRepository)

	userAuthMockData := &model.UserAuth{
		ID:         1,
		FullName:   sql.NullString{Valid: true, String: "Test Name"},
		AvatarFile: sql.NullString{Valid: true, String: "http://test.com/img.jpg"},
		Email:      "test@gmail.com",
		Username:   "05111740000184",
	}

	mockUserAuthRepository.On("GetUserAuthByID").Return(userAuthMockData, nil)

	userAuthData, err := testService.GetUserData(userAuthMockData.ID)

	mockUserAuthRepository.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.Nil(t, err)

	// Test Usecase Result is Valid
	assert.Equal(t, userAuthMockData.ID, userAuthData.ID)
	assert.Equal(t, userAuthMockData.FullName.String, userAuthData.FullName)
	assert.Equal(t, userAuthMockData.AvatarFile.String, userAuthData.AvatarFile)
	assert.Equal(t, userAuthMockData.Email, userAuthData.Email)
	assert.Equal(t, userAuthMockData.Username, userAuthData.Username)

}

func TestGetUserDataUserNotFound(t *testing.T) {

	mockUserAuthRepository := new(mysql.UserAuthRepositoryMock)

	testService := usecase.NewAccountManagerUsecase(mockUserAuthRepository)

	mockUserAuthRepository.On("GetUserAuthByID").Return(&model.UserAuth{}, fmt.Errorf("User Not Found"))

	_, err := testService.GetUserData(1)

	mockUserAuthRepository.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.NotNil(t, err)

}
