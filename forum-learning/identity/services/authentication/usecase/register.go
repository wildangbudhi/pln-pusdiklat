package usecase

import (
	"database/sql"
	"fmt"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain/event"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain/model"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil

}

func (usecase *authenticationUsecase) Register(fullName string, email domain.Email, username string, password string) (int64, error) {

	var userAuthCheck *model.UserAuth
	var err error

	userAuthCheck, err = usecase.userAuthRepository.GetUserAuthByEmail(email.GetValue())

	if userAuthCheck != nil {
		return -1, fmt.Errorf("Email Sudah Digunakan")
	}

	userAuthCheck, err = usecase.userAuthRepository.GetUserAuthByUsername(username)

	if userAuthCheck != nil {
		return -1, fmt.Errorf("Username Sudah Digunakan")
	}

	hashedPassword, err := hashPassword(password)

	if err != nil {
		return -1, fmt.Errorf("Gagal Mengenkripsi Password")
	}

	userAuth := model.UserAuth{
		FullName: sql.NullString{Valid: true, String: fullName},
		Email:    email.GetValue(),
		Username: username,
		Password: hashedPassword,
		Roles: []model.Roles{
			{ID: 1, RoleName: "Client"},
		},
	}
	userAuthID, err := usecase.userAuthRepository.InsertUserAuth(&userAuth)

	if err != nil {
		return -1, err
	}

	userAuthEvent := event.UserAuthEvent{
		Action: "CREATE",
		Data: &event.UserAuth{
			ID:       int(userAuthID),
			FullName: fullName,
			Email:    email.GetValue(),
			Roles: []event.Roles{
				{ID: 1, RoleName: "Client"},
			},
		},
	}

	err = usecase.userAuthEventRepository.PublishDataChangesEvent(&userAuthEvent)

	if err != nil {
		return -1, err
	}

	return userAuthID, nil

}
