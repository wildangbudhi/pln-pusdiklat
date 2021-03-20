package usecase

import (
	"io"
	"net/http"
)

func (usecase *authenticationUsecase) Verify(requestBody io.Reader) (*http.Response, error) {

	resp, err := usecase.authenticationRepository.Verify(requestBody)

	if err != nil {
		return nil, err
	}

	return resp, nil

}
