package usecase

import (
	"database/sql"
	"fmt"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

func (usecase *forumUsecase) UpdateForum(requestUserID int, forumID domain.UUID, title string, question string, categoriID int, requestUserRoles []string) error {

	var isAdmin, isClient bool = false, false

	for i := 0; i < len(requestUserRoles); i++ {
		if requestUserRoles[i] == "Client" {
			isClient = true
		} else if requestUserRoles[i] == "Admin" {
			isAdmin = true
		}
	}

	if isAdmin == false && isClient == false {
		return fmt.Errorf("You Don't Have Access To Update Spesific Forum Reply")
	}

	forumData, err := usecase.forumRepository.GetForumByIDWithUserReaction(forumID, requestUserID)

	if err != nil {
		return err
	}

	if !isAdmin && requestUserID != forumData.AuthorUserID {
		return fmt.Errorf("You Don't Have Access To Update Spesific Forum Reply")
	}

	rowAffected, err := usecase.forumRepository.UpdateForumByID(forumID, title, sql.NullString{String: question, Valid: true}, categoriID)

	if err != nil {
		return err
	}

	if rowAffected <= 0 {
		return fmt.Errorf("Forum Reply With Spesific ID Not Found")
	}

	return nil

}
