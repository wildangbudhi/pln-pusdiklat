package usecase

import (
	"io"
	"net/http"
)

func (usecase *authenticationUsecase) Register(requestBody io.Reader) (*http.Response, error) {

	resp, err := usecase.authenticationRepository.Register(requestBody)

	if err != nil {
		return nil, err
	}

	return resp, nil

}
