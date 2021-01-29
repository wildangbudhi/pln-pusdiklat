package usecase_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain/mocks/mysql"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/usecase"
)

func TestGetForumValid(t *testing.T) {

	forumRepositoryMock := new(mysql.ForumRepositoryMock)

	testService := usecase.NewForumUsecase(nil, nil, forumRepositoryMock, nil, nil, nil)

	requestUserID := 2
	forumID, err := domain.NewUUIDFromString("e01aa5f6-b2fb-4aae-8970-b3181f2bc8e9")

	if err != nil {
		log.Fatal(err)
	}

	mockData := &domain.Forum{
		ID: *forumID,
	}

	forumRepositoryMock.On("GetForumByIDWithUserReaction").Return(mockData, nil)

	forum, err := testService.GetForum(*forumID, requestUserID)

	forumRepositoryMock.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.Nil(t, err)

	// Test Usecase forumID same with MockData
	assert.Equal(t, forum.ID.GetValue(), mockData.ID.GetValue())

}

func TestGetForumNotFound(t *testing.T) {

	forumRepositoryMock := new(mysql.ForumRepositoryMock)

	testService := usecase.NewForumUsecase(nil, nil, forumRepositoryMock, nil, nil, nil)

	requestUserID := 2
	forumID, err := domain.NewUUIDFromString("e01aa5f6-b2fb-4aae-8970-b3181f2bc8e9")

	if err != nil {
		log.Fatal(err)
	}

	forumRepositoryMock.On("GetForumByIDWithUserReaction").Return(&domain.Forum{}, fmt.Errorf("Forum With Spesific ID Not Found"))

	_, err = testService.GetForum(*forumID, requestUserID)

	forumRepositoryMock.AssertExpectations(t)

	// Test Usecase Error is Not Nil
	assert.NotNil(t, err)

}
