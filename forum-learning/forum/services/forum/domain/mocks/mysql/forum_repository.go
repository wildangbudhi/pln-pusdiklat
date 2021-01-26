package mysql

import (
	"database/sql"

	"github.com/stretchr/testify/mock"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

type ForumRepositoryMock struct {
	mock.Mock
}

func (repo *ForumRepositoryMock) InsertForum(id domain.UUID, title string, question sql.NullString, authorUserID int, categoryID int, status string) (int, error) {

	args := repo.Called()
	result := args.Get(0)
	return result.(int), args.Error(1)

}

func (repo *ForumRepositoryMock) GetForumByIDWithUserReaction(id domain.UUID, userID int) (*domain.Forum, error) {

	args := repo.Called()
	result := args.Get(0)
	return result.(*domain.Forum), args.Error(1)

}

func (repo *ForumRepositoryMock) FetchForumWithUserReaction(offset int, limit int, userID int) ([]domain.Forum, error) {

	args := repo.Called()
	result := args.Get(0)
	return result.([]domain.Forum), args.Error(1)

}

func (repo *ForumRepositoryMock) FetchForumByAuthorIDWithUserReaction(authorID int, offset int, limit int, userID int) ([]domain.Forum, error) {

	args := repo.Called()
	result := args.Get(0)
	return result.([]domain.Forum), args.Error(1)

}

func (repo *ForumRepositoryMock) SearchByTitleAndQuestionWithUserReaction(offset int, limit int, userID int, query string) ([]domain.Forum, error) {

	args := repo.Called()
	result := args.Get(0)
	return result.([]domain.Forum), args.Error(1)

}

func (repo *ForumRepositoryMock) DeleteForumByID(id domain.UUID) (int, error) {

	args := repo.Called()
	result := args.Get(0)
	return result.(int), args.Error(1)

}

func (repo *ForumRepositoryMock) UpdateForumByID(id domain.UUID, title string, question sql.NullString, categoryID int) (int, error) {

	args := repo.Called()
	result := args.Get(0)
	return result.(int), args.Error(1)

}

func (repo *ForumRepositoryMock) UpdateForumStatusByID(id domain.UUID, status string) (int, error) {

	args := repo.Called()
	result := args.Get(0)
	return result.(int), args.Error(1)

}
