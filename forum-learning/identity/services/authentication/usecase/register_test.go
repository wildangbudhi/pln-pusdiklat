package usecase_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain/mocks/mysql"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain/mocks/rabbitmq"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain/model"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/usecase"
)

func TestRegisterSuccess(t *testing.T) {

	mockUserAuthRepository := new(mysql.UserAuthRepositoryMock)
	mockUserAuthEventRepository := new(rabbitmq.UserAuthEventRepositoryMock)

	testService := usecase.NewAuthenticationUsecase(mockUserAuthRepository, mockUserAuthEventRepository, nil)

	mockUserAuthRepository.On("GetUserAuthByEmail").Return(&model.UserAuth{}, fmt.Errorf("User Not Found"))
	mockUserAuthRepository.On("GetUserAuthByUsername").Return(&model.UserAuth{}, fmt.Errorf("User Not Found"))
	mockUserAuthRepository.On("InsertUserAuth").Return(int64(1), nil)

	mockUserAuthEventRepository.On("PublishDataChangesEvent").Return(nil)

	fullName := "Test Name"
	username := "test_username"
	password := "password123"
	email, err := domain.NewEmail("test@gmail.com")

	if err != nil {
		log.Fatal(err)
	}

	userAuthID, err := testService.Register(fullName, *email, username, password)

	mockUserAuthRepository.AssertExpectations(t)
	mockUserAuthEventRepository.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.Nil(t, err)

	// Test Usecase Response
	assert.Equal(t, int64(1), userAuthID)

}

func TestRegisterEmailAlreadyUsed(t *testing.T) {

	mockUserAuthRepository := new(mysql.UserAuthRepositoryMock)

	testService := usecase.NewAuthenticationUsecase(mockUserAuthRepository, nil, nil)

	mockUserAuthRepository.On("GetUserAuthByEmail").Return(&model.UserAuth{}, nil)

	fullName := "Test Name"
	username := "test_username"
	password := "password123"
	email, err := domain.NewEmail("test@gmail.com")

	if err != nil {
		log.Fatal(err)
	}

	_, err = testService.Register(fullName, *email, username, password)

	mockUserAuthRepository.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.NotNil(t, err)

}

func TestRegisterUsernamelAlreadyUsed(t *testing.T) {

	mockUserAuthRepository := new(mysql.UserAuthRepositoryMock)

	testService := usecase.NewAuthenticationUsecase(mockUserAuthRepository, nil, nil)

	mockUserAuthRepository.On("GetUserAuthByEmail").Return(&model.UserAuth{}, fmt.Errorf("User Not Found"))
	mockUserAuthRepository.On("GetUserAuthByUsername").Return(&model.UserAuth{}, nil)

	fullName := "Test Name"
	username := "test_username"
	password := "password123"
	email, err := domain.NewEmail("test@gmail.com")

	if err != nil {
		log.Fatal(err)
	}

	_, err = testService.Register(fullName, *email, username, password)

	mockUserAuthRepository.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.NotNil(t, err)

}

func TestRegisterFailedToInsert(t *testing.T) {

	mockUserAuthRepository := new(mysql.UserAuthRepositoryMock)

	testService := usecase.NewAuthenticationUsecase(mockUserAuthRepository, nil, nil)

	mockUserAuthRepository.On("GetUserAuthByEmail").Return(&model.UserAuth{}, fmt.Errorf("User Not Found"))
	mockUserAuthRepository.On("GetUserAuthByUsername").Return(&model.UserAuth{}, fmt.Errorf("User Not Found"))
	mockUserAuthRepository.On("InsertUserAuth").Return(int64(-1), fmt.Errorf("Failed to Insert User"))

	fullName := "Test Name"
	username := "test_username"
	password := "password123"
	email, err := domain.NewEmail("test@gmail.com")

	if err != nil {
		log.Fatal(err)
	}

	_, err = testService.Register(fullName, *email, username, password)

	mockUserAuthRepository.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.NotNil(t, err)

}

func TestRegisterFailedToPublishEvent(t *testing.T) {

	mockUserAuthRepository := new(mysql.UserAuthRepositoryMock)
	mockUserAuthEventRepository := new(rabbitmq.UserAuthEventRepositoryMock)

	testService := usecase.NewAuthenticationUsecase(mockUserAuthRepository, mockUserAuthEventRepository, nil)

	mockUserAuthRepository.On("GetUserAuthByEmail").Return(&model.UserAuth{}, fmt.Errorf("User Not Found"))
	mockUserAuthRepository.On("GetUserAuthByUsername").Return(&model.UserAuth{}, fmt.Errorf("User Not Found"))
	mockUserAuthRepository.On("InsertUserAuth").Return(int64(1), nil)

	mockUserAuthEventRepository.On("PublishDataChangesEvent").Return(fmt.Errorf("Failed to Publish Event"))

	fullName := "Test Name"
	username := "test_username"
	password := "password123"
	email, err := domain.NewEmail("test@gmail.com")

	if err != nil {
		log.Fatal(err)
	}

	_, err = testService.Register(fullName, *email, username, password)

	mockUserAuthRepository.AssertExpectations(t)
	mockUserAuthEventRepository.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.NotNil(t, err)

}
