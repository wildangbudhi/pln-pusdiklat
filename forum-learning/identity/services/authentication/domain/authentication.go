package domain

import (
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain/model"
)

// AuthenticationUsecase is an Interface for Authentication Usecase
type AuthenticationUsecase interface {
	Register(fullName string, email Email, username string, pasword string) (int64, error)
	Authenticate(username string, password string) (string, error)
	Verify(token string) (*VerifyUsecaseResponse, error)
	EndpointAuthorize(userID int, method string, url string) (bool, error)
}

type VerifyUsecaseResponse struct {
	ID       int           `json:"id"`
	FullName string        `json:"full_name"`
	Email    string        `json:"email"`
	Username string        `json:"username"`
	Roles    []model.Roles `json:"roles"`
}
