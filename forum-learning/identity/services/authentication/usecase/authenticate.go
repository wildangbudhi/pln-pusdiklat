package usecase

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func createJWTToken(userID int, expirationDate time.Time, secretKey []byte) (string, error) {

	aksesTokenClaims := jwt.MapClaims{
		"id":  strconv.Itoa(userID),
		"exp": expirationDate.Unix(),
	}

	aksesToken := jwt.NewWithClaims(jwt.SigningMethodHS512, aksesTokenClaims)

	token, err := aksesToken.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return token, nil

}

func (usecase *authenticationUsecase) Authenticate(username string, password string) (string, error) {

	userAuth, err := usecase.userAuthRepository.GetUserAuthByUsername(username)

	if err != nil {
		return "", err
	}

	if !checkPasswordHash(password, userAuth.Password) {
		return "", fmt.Errorf("Password Doesn't Match")
	}

	expirationDate := time.Now().Add(time.Hour * 24)

	token, err := createJWTToken(userAuth.ID, expirationDate, usecase.secretKey)

	if err != nil {
		return "", fmt.Errorf("Failed To Create Authentication Token")
	}

	return token, nil

}
