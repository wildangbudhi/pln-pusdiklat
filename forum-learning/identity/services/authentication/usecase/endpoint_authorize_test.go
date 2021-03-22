package usecase_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain/mocks/mysql"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain/model"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/usecase"
)

func TestEnpointAuthorizeAuthorized(t *testing.T) {

	mockUserAuthRepository := new(mysql.UserAuthRepositoryMock)

	testService := usecase.NewAuthenticationUsecase(mockUserAuthRepository, nil, nil)

	userAuthMockData := &model.UserAuth{
		Roles: []model.Roles{
			{1, "Client"},
		},
	}

	mockUserAuthRepository.On("GetUserAuthByID").Return(userAuthMockData, nil)
	mockUserAuthRepository.On("CountRoleActivitiesPermission").Return(1, nil)

	authorized, err := testService.EndpointAuthorize(1, "GET", "/auth/")

	mockUserAuthRepository.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.Nil(t, err)

	// Test for Usecase Result
	assert.Equal(t, true, authorized)

}

func TestEnpointAuthorizeUnauthorized(t *testing.T) {

	mockUserAuthRepository := new(mysql.UserAuthRepositoryMock)

	testService := usecase.NewAuthenticationUsecase(mockUserAuthRepository, nil, nil)

	userAuthMockData := &model.UserAuth{
		Roles: []model.Roles{
			{1, "Client"},
		},
	}

	mockUserAuthRepository.On("GetUserAuthByID").Return(userAuthMockData, nil)
	mockUserAuthRepository.On("CountRoleActivitiesPermission").Return(0, nil)

	authorized, err := testService.EndpointAuthorize(1, "GET", "/auth/")

	mockUserAuthRepository.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.Nil(t, err)

	// Test for Usecase Result
	assert.Equal(t, false, authorized)

}

func TestEnpointAuthorizeUserNotFound(t *testing.T) {

	mockUserAuthRepository := new(mysql.UserAuthRepositoryMock)

	testService := usecase.NewAuthenticationUsecase(mockUserAuthRepository, nil, nil)

	mockUserAuthRepository.On("GetUserAuthByID").Return(&model.UserAuth{}, fmt.Errorf("User Not Found"))

	_, err := testService.EndpointAuthorize(1, "GET", "/auth/")

	mockUserAuthRepository.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.NotNil(t, err)

}
