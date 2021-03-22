package usecase

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain/model"
)

type authenticationUsecase struct {
	userAuthRepository model.UserAuthRepository
	secretKey          []byte
	apiSecretKey       []byte
}

// NewAuthenticationUsecase is a Constructor of authenticationUsecase
// Which implement AuthenticationUsecase Interface
func NewAuthenticationUsecase(userAuthRepository model.UserAuthRepository, secretKey []byte, apiSecretKey []byte) domain.AuthenticationUsecase {
	return &authenticationUsecase{
		userAuthRepository: userAuthRepository,
		secretKey:          secretKey,
		apiSecretKey:       apiSecretKey,
	}
}

func (usercase *authenticationUsecase) VerifyToken(stringToken string, secretKey []byte) (*jwt.Token, error) {
	token, err := jwt.Parse(stringToken, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected Signing Method")
		}

		return secretKey, nil

	})

	if err != nil {
		return nil, err
	}

	return token, nil

}
