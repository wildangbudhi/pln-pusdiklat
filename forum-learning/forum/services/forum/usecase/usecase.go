package usecase

import "github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"

type forumUsecase struct {
	userAuthRepository             domain.UserAuthRepository
	categoryRepository             domain.CategoryRepository
	forumRepository                domain.ForumRepository
	forumReactionRepository        domain.ForumReactionRepository
	forumRepliesRepository         domain.ForumRepliesRepository
	forumRepliesReactionRepository domain.ForumRepliesReactionRepository
}

func NewForumUsecase(
	userAuthRepository domain.UserAuthRepository,
	categoryRepository domain.CategoryRepository,
	forumRepository domain.ForumRepository,
	forumReactionRepository domain.ForumReactionRepository,
	forumRepliesRepository domain.ForumRepliesRepository,
	forumRepliesReactionRepository domain.ForumRepliesReactionRepository,
) domain.ForumUsecase {

	return &forumUsecase{
		userAuthRepository:             userAuthRepository,
		categoryRepository:             categoryRepository,
		forumRepository:                forumRepository,
		forumReactionRepository:        forumReactionRepository,
		forumRepliesRepository:         forumRepliesRepository,
		forumRepliesReactionRepository: forumRepliesReactionRepository,
	}

}
