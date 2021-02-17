package usecase

import (
	"fmt"
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

	return rowAffected > 0, nil

}
