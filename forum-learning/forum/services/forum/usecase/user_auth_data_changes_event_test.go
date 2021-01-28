package usecase_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain/mocks/mysql"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/usecase"
)

func TestUserAuthDataChangesEventEventNotAllowed(t *testing.T) {

	testService := usecase.NewForumUsecase(nil, nil, nil, nil, nil, nil)

	eventAction := "PUT"
	userID := 1
	fullName := "Test Name"
	email, err := domain.NewEmail("test@gmail.com")
	username := "05111740000184"

	err = testService.UserAuthDataChangesEvent(eventAction, userID, fullName, *email, username)

	// Test Usecase Error is Nil
	assert.NotNil(t, err)

}

func TestUserAuthDataChangesEventCreateSuccess(t *testing.T) {

	userAuthRepositoryMock := new(mysql.UserAuthRepositoryMock)

	testService := usecase.NewForumUsecase(userAuthRepositoryMock, nil, nil, nil, nil, nil)

	eventAction := "CREATE"
	userID := 1
	fullName := "Test Name"
	email, _ := domain.NewEmail("test@gmail.com")
	username := "05111740000184"

	userAuthRepositoryMock.On("InsertUserAuth").Return(int64(1), nil)

	err := testService.UserAuthDataChangesEvent(eventAction, userID, fullName, *email, username)

	userAuthRepositoryMock.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.Nil(t, err)

}

func TestUserAuthDataChangesEventCreateError(t *testing.T) {

	userAuthRepositoryMock := new(mysql.UserAuthRepositoryMock)

	testService := usecase.NewForumUsecase(userAuthRepositoryMock, nil, nil, nil, nil, nil)

	eventAction := "CREATE"
	userID := 1
	fullName := "Test Name"
	email, _ := domain.NewEmail("test@gmail.com")
	username := "05111740000184"

	userAuthRepositoryMock.On("InsertUserAuth").Return(int64(-1), fmt.Errorf("Error Inster User Auth Data"))

	err := testService.UserAuthDataChangesEvent(eventAction, userID, fullName, *email, username)

	userAuthRepositoryMock.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.NotNil(t, err)

}

func TestUserAuthDataChangesEventUpdateSuccess(t *testing.T) {

	userAuthRepositoryMock := new(mysql.UserAuthRepositoryMock)

	testService := usecase.NewForumUsecase(userAuthRepositoryMock, nil, nil, nil, nil, nil)

	eventAction := "UPDATE"
	userID := 1
	fullName := "Test Name"
	email, _ := domain.NewEmail("test@gmail.com")
	username := "05111740000184"

	userAuthRepositoryMock.On("UpdateUserAuthByID").Return(1, nil)

	err := testService.UserAuthDataChangesEvent(eventAction, userID, fullName, *email, username)

	userAuthRepositoryMock.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.Nil(t, err)

}

func TestUserAuthDataChangesEventUpdateError(t *testing.T) {

	userAuthRepositoryMock := new(mysql.UserAuthRepositoryMock)

	testService := usecase.NewForumUsecase(userAuthRepositoryMock, nil, nil, nil, nil, nil)

	eventAction := "UPDATE"
	userID := 1
	fullName := "Test Name"
	email, _ := domain.NewEmail("test@gmail.com")
	username := "05111740000184"

	userAuthRepositoryMock.On("UpdateUserAuthByID").Return(-1, fmt.Errorf("Error Updating User Auth Data"))

	err := testService.UserAuthDataChangesEvent(eventAction, userID, fullName, *email, username)

	userAuthRepositoryMock.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.NotNil(t, err)

}
