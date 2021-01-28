package usecase

import "github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"

func (usecase *forumUsecase) ReactForumReplies(requestUserID int, forumRepliesID domain.UUID, userReaction domain.ForumRepliesReactionType) error {

	forumRepliesReaction, err := usecase.forumRepliesReactionRepository.GetForumRepliesReactionByUserIDAndForumRepliesID(requestUserID, forumRepliesID)

	if err != nil || forumRepliesReaction == nil {

		forumRepliesReaction = &domain.ForumRepliesReaction{
			UserID:         requestUserID,
			ForumRepliesID: forumRepliesID,
			UpVote:         userReaction.IsUpVoteToggled(),
			DownVote:       userReaction.IsDownVoteToggled(),
			Agree:          userReaction.IsAgreeToggled(),
			Skeptic:        userReaction.IsSkepticToggled(),
		}

		_, err = usecase.forumRepliesReactionRepository.InsertForumRepliesReaction(forumRepliesReaction)

	} else {
		_, err = usecase.forumRepliesReactionRepository.UpdateForumRepliesReactionByUserIDAndForumRepliesID(requestUserID, forumRepliesID, userReaction)
	}

	return err

}
