package usecase

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (usecase *accountManagerUsecase) UpdateUserData(authenticationToken string, userID int, requestBody io.Reader, requestURI string, requestMethod string) (*http.Response, error) {

	verifyResponse, err := usecase.Authenticate(authenticationToken)

	if err != nil {
		return nil, err
	}

	requestMethod = strings.ToUpper(requestMethod)

	authorizeResponse, err := usecase.Authorize(verifyResponse.ID, requestMethod, requestURI)

	if err != nil {
		return nil, err
	}

	if !authorizeResponse.Authorized {
		return nil, fmt.Errorf("You Dont Have Access To This Resource")
	}

	xAuthRoles := []string{}

	for i := 0; i < len(verifyResponse.Roles); i++ {
		xAuthRoles = append(xAuthRoles, verifyResponse.Roles[i].RoleName)
	}

	res, err := usecase.accountManagerRepository.UpdateUserData(userID, requestBody, verifyResponse.ID, xAuthRoles)

	if err != nil {
		return nil, err
	}

	return res, nil

}
