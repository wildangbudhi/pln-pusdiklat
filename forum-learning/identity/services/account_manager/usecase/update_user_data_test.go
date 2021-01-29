package usecase_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/account_manager/domain/mocks/mysql"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/account_manager/domain/mocks/rabbitmq"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/account_manager/domain/model"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/account_manager/usecase"
)

func TestUpdateUserDataSuccessClient(t *testing.T) {

	mockUserAuthRepository := new(mysql.UserAuthRepositoryMock)
	mockUserAuthEventRepository := new(rabbitmq.UserAuthEventRepositoryMock)

	testService := usecase.NewAccountManagerUsecase(mockUserAuthRepository, mockUserAuthEventRepository)

	userAuthMockData := &model.UserAuth{
		ID:         1,
		FullName:   sql.NullString{Valid: true, String: "Test Name"},
		AvatarFile: sql.NullString{Valid: true, String: "http://test.com/img.jpg"},
		Email:      "test@gmail.com",
		Username:   "05111740000184",
	}

	mockUserAuthRepository.On("UpdateUserAuthByID").Return(1, nil)
	mockUserAuthRepository.On("GetUserAuthByID").Return(userAuthMockData, nil)
	mockUserAuthEventRepository.On("PublishDataChangesEvent").Return(nil)

	status, err := testService.UpdateUserData(1, "Test Name", 1, []string{"Client"})

	mockUserAuthRepository.AssertExpectations(t)
	mockUserAuthEventRepository.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.Nil(t, err)

	// Test Usecase Result is Valid
	assert.Equal(t, true, status)

}

func TestUpdateUserDataSuccessAdmin(t *testing.T) {

	mockUserAuthRepository := new(mysql.UserAuthRepositoryMock)
	mockUserAuthEventRepository := new(rabbitmq.UserAuthEventRepositoryMock)

	testService := usecase.NewAccountManagerUsecase(mockUserAuthRepository, mockUserAuthEventRepository)

	userAuthMockData := &model.UserAuth{
		ID:         1,
		FullName:   sql.NullString{Valid: true, String: "Test Name"},
		AvatarFile: sql.NullString{Valid: true, String: "http://test.com/img.jpg"},
		Email:      "test@gmail.com",
		Username:   "05111740000184",
	}

	mockUserAuthRepository.On("UpdateUserAuthByID").Return(1, nil)
	mockUserAuthRepository.On("GetUserAuthByID").Return(userAuthMockData, nil)
	mockUserAuthEventRepository.On("PublishDataChangesEvent").Return(nil)

	status, err := testService.UpdateUserData(1, "Test Name", 2, []string{"Client", "Admin"})

	mockUserAuthRepository.AssertExpectations(t)
	mockUserAuthEventRepository.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.Nil(t, err)

	// Test Usecase Result is Valid
	assert.Equal(t, true, status)

}

func TestUpdateUserDataUpdateOtherNotAdmin(t *testing.T) {

	testService := usecase.NewAccountManagerUsecase(nil, nil)

	_, err := testService.UpdateUserData(1, "Test Name", 2, []string{"Client"})

	// Test Usecase Error is Nil
	assert.NotNil(t, err)

}

func TestUpdateUserDataFailedToUpdate(t *testing.T) {

	mockUserAuthRepository := new(mysql.UserAuthRepositoryMock)
	mockUserAuthEventRepository := new(rabbitmq.UserAuthEventRepositoryMock)

	testService := usecase.NewAccountManagerUsecase(mockUserAuthRepository, mockUserAuthEventRepository)

	userAuthMockData := &model.UserAuth{
		ID:         1,
		FullName:   sql.NullString{Valid: true, String: "Test Name"},
		AvatarFile: sql.NullString{Valid: true, String: "http://test.com/img.jpg"},
		Email:      "test@gmail.com",
		Username:   "05111740000184",
	}

	mockUserAuthRepository.On("UpdateUserAuthByID").Return(0, nil)
	mockUserAuthRepository.On("GetUserAuthByID").Return(userAuthMockData, nil)
	mockUserAuthEventRepository.On("PublishDataChangesEvent").Return(nil)

	status, err := testService.UpdateUserData(1, "Test Name", 1, []string{"Client"})

	mockUserAuthRepository.AssertExpectations(t)
	mockUserAuthEventRepository.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.Nil(t, err)

	// Test Usecase Result is Valid
	assert.Equal(t, false, status)

}

func TestUpdateUserDataFailedToPublishEvent(t *testing.T) {

	mockUserAuthRepository := new(mysql.UserAuthRepositoryMock)
	mockUserAuthEventRepository := new(rabbitmq.UserAuthEventRepositoryMock)

	testService := usecase.NewAccountManagerUsecase(mockUserAuthRepository, mockUserAuthEventRepository)

	userAuthMockData := &model.UserAuth{
		ID:         1,
		FullName:   sql.NullString{Valid: true, String: "Test Name"},
		AvatarFile: sql.NullString{Valid: true, String: "http://test.com/img.jpg"},
		Email:      "test@gmail.com",
		Username:   "05111740000184",
	}

	mockUserAuthRepository.On("UpdateUserAuthByID").Return(0, nil)
	mockUserAuthRepository.On("GetUserAuthByID").Return(userAuthMockData, nil)
	mockUserAuthEventRepository.On("PublishDataChangesEvent").Return(fmt.Errorf("Failed To Publish Event"))

	_, err := testService.UpdateUserData(1, "Test Name", 1, []string{"Client"})

	mockUserAuthRepository.AssertExpectations(t)
	mockUserAuthEventRepository.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.NotNil(t, err)

}
