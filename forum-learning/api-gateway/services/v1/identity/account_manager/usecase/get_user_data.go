package usecase

import (
	"fmt"
	"net/http"
	"strings"
)

func (usecase *accountManagerUsecase) GetUserData(authenticationToken string, userID int, requestURI string, requestMethod string) (*http.Response, error) {

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

	res, err := usecase.accountManagerRepository.GetUserData(userID)

	if err != nil {
		return nil, err
	}

	return res, nil

}
