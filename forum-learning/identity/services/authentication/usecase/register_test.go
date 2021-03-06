package usecase_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain/mocks/mysql"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain/model"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/usecase"
)

func TestRegisterSuccess(t *testing.T) {

	mockUserAuthRepository := new(mysql.UserAuthRepositoryMock)

	testService := usecase.NewAuthenticationUsecase(mockUserAuthRepository, nil, nil)

	mockUserAuthRepository.On("GetUserAuthByUsername").Return(&model.UserAuth{}, fmt.Errorf("User Not Found"))
	mockUserAuthRepository.On("InsertUserAuth").Return(int64(1), nil)

	fullName := "Test Name"
	username := "test_username"
	password := "password123"

	userAuthID, err := testService.Register(fullName, username, password)

	mockUserAuthRepository.AssertExpectations(t)
	// mockUserAuthEventRepository.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.Nil(t, err)

	// Test Usecase Response
	assert.Equal(t, int64(1), userAuthID)

}

func TestRegisterUsernamelAlreadyUsed(t *testing.T) {

	mockUserAuthRepository := new(mysql.UserAuthRepositoryMock)

	testService := usecase.NewAuthenticationUsecase(mockUserAuthRepository, nil, nil)

	mockUserAuthRepository.On("GetUserAuthByUsername").Return(&model.UserAuth{}, nil)

	fullName := "Test Name"
	username := "test_username"
	password := "password123"

	_, err := testService.Register(fullName, username, password)

	mockUserAuthRepository.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.NotNil(t, err)

}

func TestRegisterFailedToInsert(t *testing.T) {

	mockUserAuthRepository := new(mysql.UserAuthRepositoryMock)

	testService := usecase.NewAuthenticationUsecase(mockUserAuthRepository, nil, nil)

	mockUserAuthRepository.On("GetUserAuthByUsername").Return(&model.UserAuth{}, fmt.Errorf("User Not Found"))
	mockUserAuthRepository.On("InsertUserAuth").Return(int64(-1), fmt.Errorf("Failed to Insert User"))

	fullName := "Test Name"
	username := "test_username"
	password := "password123"

	_, err := testService.Register(fullName, username, password)

	mockUserAuthRepository.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.NotNil(t, err)

}
