package domain

import (
	"io"
	"net/http"
)

type AuthenticationRepository interface {
	// Register(requestBody io.Reader) (*http.Response, error)
	Authenticate(requestBody io.Reader) (*http.Response, error)
}
