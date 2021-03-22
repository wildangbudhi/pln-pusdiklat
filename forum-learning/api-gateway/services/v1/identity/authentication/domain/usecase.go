package domain

import (
	"io"
	"net/http"
)

type AuthenticationUsecase interface {
	Register(requestBody io.Reader) (*http.Response, error)
	Authenticate(requestBody io.Reader) (*http.Response, error)
	Verify(requestBody io.Reader) (*http.Response, error)
	VerifySSOPLN(requestBody io.Reader) (*http.Response, error)
	EndpointAuthorize(requestBody io.Reader) (*http.Response, error)
}
