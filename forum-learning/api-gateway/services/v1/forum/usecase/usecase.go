package usecase

import "github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/forum/domain"

type forumUsecase struct {
	forumRepository          domain.ForumRepository
	authenticationRepository domain.AuthenticationRepository
}

// NewForumUsecase is a Constructor of forumUsecase
// Which implement ForumUsecase Interface
func NewForumUsecase(forumRepository domain.ForumRepository, authenticationRepository domain.AuthenticationRepository) domain.ForumUsecase {
	return &forumUsecase{
		forumRepository:          forumRepository,
		authenticationRepository: authenticationRepository,
	}
}
