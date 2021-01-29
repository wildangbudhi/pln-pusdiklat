package usecase

import "github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"

func (usecase *forumUsecase) ReactForum(requestUserID int, forumID domain.UUID, userReaction domain.ForumReactionType) error {

	forumReaction, err := usecase.forumReactionRepository.GetForumReactionByUserIDAndForumID(requestUserID, forumID)

	if err != nil || forumReaction == nil {

		forumReaction = &domain.ForumReaction{
			UserID:   requestUserID,
			ForumID:  forumID,
			UpVote:   userReaction.IsUpVoteToggled(),
			DownVote: userReaction.IsDownVoteToggled(),
		}

		_, err = usecase.forumReactionRepository.InsertForumReaction(forumReaction)

	} else {
		_, err = usecase.forumReactionRepository.UpdateForumReactionByUserIDAndForumID(requestUserID, forumID, userReaction)
	}

	return err

}
