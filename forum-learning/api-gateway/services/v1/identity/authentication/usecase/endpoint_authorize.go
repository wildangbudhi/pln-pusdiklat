package usecase

import (
	"io"
	"net/http"
)

func (usecase *authenticationUsecase) EndpointAuthorize(requestBody io.Reader) (*http.Response, error) {

	resp, err := usecase.authenticationRepository.EndpointAuthorize(requestBody)

	if err != nil {
		return nil, err
	}

	return resp, nil

}
