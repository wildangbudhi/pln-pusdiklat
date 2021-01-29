package usecase

import (
	"fmt"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/account_manager/domain/event"
)

func (usecase *accountManagerUsecase) UpdateUserData(userID int, fullName string, requestUserID int, requestUserRoles []string) (bool, error) {

	isAdmin := false

	for i := 0; i < len(requestUserRoles); i++ {
		if requestUserRoles[i] == "Admin" {
			isAdmin = true
			break
		}
	}

	if !isAdmin && userID != requestUserID {
		return false, fmt.Errorf("You Dont Have Permission")
	}

	rowAffected, err := usecase.userAuthRepository.UpdateUserAuthByID(userID, fullName)

	if err != nil {
		return false, err
	}

	userAuthData, err := usecase.userAuthRepository.GetUserAuthByID(userID)

	if err != nil {
		return false, err
	}

	event := &event.UserAuthEvent{
		Action: "UPDATE",
		Data: &event.UserAuth{
			ID:       userID,
			FullName: fullName,
			Email:    userAuthData.Email,
			Username: userAuthData.Username,
		},
	}

	err = usecase.userAuthEventRepository.PublishDataChangesEvent(event)

	if err != nil {
		return false, err
	}

	return rowAffected > 0, nil

}
