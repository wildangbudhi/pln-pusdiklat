package domain

import (
	"io"
	"net/http"
)

type AccountManagerUsecase interface {
	GetUserData(authenticationToken string, userID int, requestURI string, requestMethod string) (*http.Response, error)
	UpdateUserData(authenticationToken string, userID int, requestBody io.Reader, requestURI string, requestMethod string) (*http.Response, error)
}
