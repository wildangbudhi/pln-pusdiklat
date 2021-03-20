package usecase

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (usecase *forumUsecase) CloseForum(requestURI string, requestMethod string, forumID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error) {

	authenticationToken, ok := requestHeader["Authorization"]

	if !ok || len(authenticationToken) == 0 {
		return nil, fmt.Errorf("Authorization Header Not Found")
	}

	delete(requestHeader, "Authorization")

	verifyResponse, err := usecase.Authenticate(authenticationToken[0])

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

	err = usecase.setAuthHeader(verifyResponse, requestHeader)

	if err != nil {
		return nil, err
	}

	resp, err := usecase.forumRepository.CloseForum(forumID, requestBody, urlQuery, requestHeader)

	if err != nil {
		return nil, err
	}

	return resp, nil

}
