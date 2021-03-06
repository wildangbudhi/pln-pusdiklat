package usecase

import (
	"io"
	"net/http"
)

func (usecase *authenticationUsecase) Authenticate(requestBody io.Reader) (*http.Response, error) {

	resp, err := usecase.authenticationRepository.Authenticate(requestBody)

	if err != nil {
		return nil, err
	}

	return resp, nil

}
