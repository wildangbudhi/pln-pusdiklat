package usecase

import (
	"database/sql"
	"fmt"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/event_consumer/domain"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/event_consumer/domain/model"
)

func (usecase *eventConsumerUsecase) UserAuthDataChangesEvent(eventAction string, id int, fullName string, email domain.Email, username string) error {

	if eventAction != "CREATE" && eventAction != "UPDATE" {
		return fmt.Errorf("Event Action Not Allowed")
	}

	var err error = nil

	if eventAction == "CREATE" {

		userAuthData := model.UserAuth{
			ID:         id,
			FullName:   sql.NullString{String: fullName, Valid: true},
			AvatarFile: sql.NullString{String: "", Valid: false},
			Email:      email.GetValue(),
			Username:   username,
		}

		// var insertedID int64

		_, err = usecase.userAuthRepository.InsertUserAuth(&userAuthData)

		// if insertedID < 0 {
		// 	err = fmt.Errorf("Failed to Insert New User Auth Data")
		// }

	} else if eventAction == "UPDATE" {

		// var rowAffected int
		_, err = usecase.userAuthRepository.UpdateUserAuthByID(id, fullName)

		// if rowAffected <= 0 {
		// 	err = fmt.Errorf("Failed to Update User Auth Data")
		// }

	}

	return err

}
