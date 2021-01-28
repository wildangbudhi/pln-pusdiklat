package usecase

import (
	"fmt"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

func (usecase *forumUsecase) DeleteForumReplies(requestUserID int, requestUserRoles []string, forumRepliesID domain.UUID) error {

	var isAdmin, isClient bool = false, false

	for i := 0; i < len(requestUserRoles); i++ {
		if requestUserRoles[i] == "Client" {
			isClient = true
		} else if requestUserRoles[i] == "Admin" {
			isAdmin = true
		}
	}

	if isAdmin == false && isClient == false {
		return fmt.Errorf("You Don't Have Access To Delete Forum")
	}

	forumRepliesData, err := usecase.forumRepliesRepository.GetForumRepliesByIDWithUserReaction(forumRepliesID, requestUserID)

	if err != nil {
		return err
	}

	if !isAdmin && requestUserID != forumRepliesData.AuthorUserID {
		return fmt.Errorf("You Don't Have Access To Delete Forum")
	}

	rowAffected, err := usecase.forumRepliesRepository.DeleteForumRepliesByID(forumRepliesID)

	if err != nil {
		return err
	}

	if rowAffected <= 0 {
		return fmt.Errorf("Forum With Spesific ID Not Found")
	}

	return nil

}
