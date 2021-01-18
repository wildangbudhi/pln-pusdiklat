package usecase

import (
	"fmt"
	"strings"
)

func (usecase *authenticationUsecase) EndpointAuthorize(userID int, method string, url string) (bool, error) {

	method = strings.ToUpper(method)

	if method != "GET" && method != "POST" {
		return false, fmt.Errorf("Method Tidak Izinkan")
	}

	userAuth, err := usecase.userAuthRepository.GetUserAuthByID(userID)

	if err != nil {
		return false, err
	}

	userRoleIDList := make([]int, 0)

	for i := 0; i < len(userAuth.Roles); i++ {
		userRoleIDList = append(userRoleIDList, userAuth.Roles[i].ID)
	}

	permisionCount, err := usecase.userAuthRepository.CountRoleActivitiesPermission(method, url, userRoleIDList)

	if err != nil {
		return false, err
	}

	return permisionCount > 0, nil

}
