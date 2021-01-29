package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain/mocks/mysql"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/usecase"
)

func TestCreateForumValidTitleEmpty(t *testing.T) {

	testService := usecase.NewForumUsecase(nil, nil, nil, nil, nil, nil)

	title := ""
	question := "Saya sudah mencoba nya"
	requestUserID := 1
	categoryID := 18

	_, err := testService.CreateForum(title, question, requestUserID, categoryID)

	// Test Usecase Error is Nil
	assert.NotNil(t, err)

}

func TestCreateForumValid(t *testing.T) {

	forumRepositoryMock := new(mysql.ForumRepositoryMock)

	testService := usecase.NewForumUsecase(nil, nil, forumRepositoryMock, nil, nil, nil)

	title := "Test"
	question := "Saya sudah mencoba nya"
	requestUserID := 1
	categoryID := 18

	forumRepositoryMock.On("InsertForum").Return(1, nil)

	_, err := testService.CreateForum(title, question, requestUserID, categoryID)

	forumRepositoryMock.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.Nil(t, err)

}
