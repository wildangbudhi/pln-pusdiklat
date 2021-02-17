package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/account_manager/domain/mocks/mysql"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/account_manager/usecase"
)

func TestUpdateUserDataSuccessClient(t *testing.T) {

	mockUserAuthRepository := new(mysql.UserAuthRepositoryMock)

	testService := usecase.NewAccountManagerUsecase(mockUserAuthRepository)

	mockUserAuthRepository.On("UpdateUserAuthByID").Return(1, nil)

	status, err := testService.UpdateUserData(1, "Test Name", 1, []string{"Client"})

	mockUserAuthRepository.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.Nil(t, err)

	// Test Usecase Result is Valid
	assert.Equal(t, true, status)

}

func TestUpdateUserDataSuccessAdmin(t *testing.T) {

	mockUserAuthRepository := new(mysql.UserAuthRepositoryMock)

	testService := usecase.NewAccountManagerUsecase(mockUserAuthRepository)

	mockUserAuthRepository.On("UpdateUserAuthByID").Return(1, nil)

	status, err := testService.UpdateUserData(1, "Test Name", 2, []string{"Client", "Admin"})

	mockUserAuthRepository.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.Nil(t, err)

	// Test Usecase Result is Valid
	assert.Equal(t, true, status)

}

func TestUpdateUserDataUpdateOtherNotAdmin(t *testing.T) {

	testService := usecase.NewAccountManagerUsecase(nil)

	_, err := testService.UpdateUserData(1, "Test Name", 2, []string{"Client"})

	// Test Usecase Error is Nil
	assert.NotNil(t, err)

}

func TestUpdateUserDataFailedToUpdate(t *testing.T) {

	mockUserAuthRepository := new(mysql.UserAuthRepositoryMock)

	testService := usecase.NewAccountManagerUsecase(mockUserAuthRepository)

	mockUserAuthRepository.On("UpdateUserAuthByID").Return(0, nil)

	status, err := testService.UpdateUserData(1, "Test Name", 1, []string{"Client"})

	mockUserAuthRepository.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.Nil(t, err)

	// Test Usecase Result is Valid
	assert.Equal(t, false, status)

}
