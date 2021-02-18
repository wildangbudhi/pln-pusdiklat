package domain

import (
	"io"
	"net/http"
)

type Roles struct {
	ID       int    `json:"id" binding:"required"`
	RoleName string `json:"role_name" binding:"required"`
}

type VerifyResponse struct {
	ID       int     `json:"id" binding:"required"`
	FullName string  `json:"full_name" binding:"required"`
	Email    string  `json:"email" binding:"required"`
	Username string  `json:"username" binding:"required"`
	Roles    []Roles `json:"roles" binding:"required"`
}

type EndpointAuthorizeResponse struct {
	Authorized bool `json:"authorized" binding:"required"`
}

type EndpointAuthorizeParameter struct {
	UserID int    `json:"user_id" binding:"required"`
	Method string `json:"method" binding:"required"`
	URL    string `json:"url" binding:"required"`
}

type AuthenticationRepository interface {
	Verify(token string) (*VerifyResponse, error)
	EndpointAuthorize(data *EndpointAuthorizeParameter) (*EndpointAuthorizeResponse, error)
}

type AccountManagerRepository interface {
	GetUserData(userID int) (*http.Response, error)
	UpdateUserData(userID int, requestBody io.Reader, xAuthID int, xAuthRoles []string) (*http.Response, error)
}
