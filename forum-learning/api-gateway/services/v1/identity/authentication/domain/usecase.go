package domain

import (
	"io"
	"net/http"
)

type AuthenticationUsecase interface {
	Authenticate(requestBody io.Reader) (*http.Response, error)
}
