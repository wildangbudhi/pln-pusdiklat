package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain/mocks/mysql"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/usecase"
)

func TestFetchCategory(t *testing.T) {

	categoryRepositoryMock := new(mysql.CategoryRepositoryMock)

	testService := usecase.NewForumUsecase(nil, categoryRepositoryMock, nil, nil, nil, nil)

	mockData := []domain.Category{
		{
			ID:           18,
			CategoryName: "Bebas",
		},
	}

	categoryRepositoryMock.On("FetchCategory").Return(mockData, nil)

	categoryList, err := testService.FetchCategory()

	categoryRepositoryMock.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.Nil(t, err)

	// Test CategoryList is The Same
	assert.Equal(t, categoryList, mockData)

}
