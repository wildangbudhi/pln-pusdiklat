package usecase

import (
	"database/sql"
	"fmt"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

func (usecase *forumUsecase) UserAuthDataChangesEvent(eventAction string, id int, fullName string, email domain.Email, username string) error {

	if eventAction != "CREATE" && eventAction != "UPDATE" {
		return fmt.Errorf("Event Action Not Allowed")
	}

	var err error = nil

	if eventAction == "CREATE" {

		userAuthData := domain.UserAuth{
			ID:         id,
			FullName:   sql.NullString{String: fullName, Valid: true},
			AvatarFile: sql.NullString{String: "", Valid: false},
			Email:      email,
			Username:   username,
		}

		_, err = usecase.userAuthRepository.InsertUserAuth(&userAuthData)

	} else if eventAction == "UPDATE" {

		_, err = usecase.userAuthRepository.UpdateUserAuthByID(id, fullName)

	}

	return err

}
