package usecase

import (
	"io"
	"net/http"
)

func (usecase *authenticationUsecase) VerifySSOPLN(requestBody io.Reader) (*http.Response, error) {

	resp, err := usecase.authenticationRepository.VerifySSOPLN(requestBody)

	if err != nil {
		return nil, err
	}

	return resp, nil

}
