package usecase_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain/mocks/mysql"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain/model"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/usecase"
)

func getTokenData(stringToken string, secretKey []byte) (jwt.MapClaims, error) {
	token, err := jwt.Parse(stringToken, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil

	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, fmt.Errorf("Token Tidak Valid")
	}

	return claims, nil
}

func TestAuthenticateValid(t *testing.T) {

	secretKey := []byte("\xec\xbb\x81\x1fy\xff\tDi\xca\xc9\xd5\x92f{L\xadNh}fz\xe5\x04HS\x92x\x1f\xf0\xd2c,\xb0\xf2Z\xcfz\ru\x86\xfb)%\x89\xc5\x89Im\x84\xde\xeb\x15\xe6\xe5\x04A\xa5p\xeal\x97\xcb\xb7<\xb8y\xfb\xa0;V h\x0f\xc0YK\r\xa3\x8cq\x9f\x19?\xdf\n\xd8B\r \xe7s-\xd1\x1dG\x1bw\xa1\xef\x8f\xc6\xbe\x98\x90\xa7\xf4g\xc1\xcfn@\xe2\x83\x8b\xfb\xbb+\x94d\xb3\x98fD\x87\xe9\xe6m\x99\xee&_\xf9\xd1p\x99\xe7\x99}\xd9\x1b\x1fIj\x836r\xad\xff\xfd\x8dt\xcdFe\x9c\x8c\xd5S\x8a\xe2U\xad\xbd\xccw\xe6\xaf\xec\x0c\xd54?X\xf1\x15\xf1i\x01\x9er\x120\xb8\x05}~\x92BY\x14\xf1\xf5R\n|\xa5\xf7'\xbb\xe5,\x84\xbf\xe8\x0eH\xc3\x9b`\xc0u\xedj\x10Y\xb7\xcbu\xcf:\x8d\x93\xd6\xd0\xe3z)W*z\xd6\xc6\xb6\xd2'\xbfD\x16`]\x12\xcb\x7f[\xfc\xd0\xed\x869o\xa0\xef\xe0\xa3\xa0")

	mockUserAuthRepository := new(mysql.UserAuthRepositoryMock)

	testService := usecase.NewAuthenticationUsecase(mockUserAuthRepository, secretKey, nil)

	userAuthMockData := &model.UserAuth{
		ID:       1,
		Username: "05111740000184",
		Password: "$2a$14$JC4.1C0npGNHT8E03/O54.Clq5a/pAthGEnI01wbYFWU8p7KPqaG2",
	}

	mockUserAuthRepository.On("GetUserAuthByUsername").Return(userAuthMockData, nil)

	username := "05111740000184"
	password := "password123"

	token, err := testService.Authenticate(username, password)

	mockUserAuthRepository.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.Nil(t, err)

	jwtTokenClaims, err := getTokenData(token, secretKey)

	// Test Token Decoding Error is Nil
	assert.Nil(t, err)

	userID, ok := jwtTokenClaims["id"].(string)

	// Test if ID in JWT Token Data
	assert.Equal(t, true, ok)

	userIDInt, err := strconv.Atoi(userID)

	// Test if ID able to convert to INT
	assert.Equal(t, true, ok)

	// Test if ID same with Mock Data
	assert.Equal(t, userAuthMockData.ID, userIDInt)

}

func TestAuthenticatePasswordInValid(t *testing.T) {

	secretKey := []byte("\xec\xbb\x81\x1fy\xff\tDi\xca\xc9\xd5\x92f{L\xadNh}fz\xe5\x04HS\x92x\x1f\xf0\xd2c,\xb0\xf2Z\xcfz\ru\x86\xfb)%\x89\xc5\x89Im\x84\xde\xeb\x15\xe6\xe5\x04A\xa5p\xeal\x97\xcb\xb7<\xb8y\xfb\xa0;V h\x0f\xc0YK\r\xa3\x8cq\x9f\x19?\xdf\n\xd8B\r \xe7s-\xd1\x1dG\x1bw\xa1\xef\x8f\xc6\xbe\x98\x90\xa7\xf4g\xc1\xcfn@\xe2\x83\x8b\xfb\xbb+\x94d\xb3\x98fD\x87\xe9\xe6m\x99\xee&_\xf9\xd1p\x99\xe7\x99}\xd9\x1b\x1fIj\x836r\xad\xff\xfd\x8dt\xcdFe\x9c\x8c\xd5S\x8a\xe2U\xad\xbd\xccw\xe6\xaf\xec\x0c\xd54?X\xf1\x15\xf1i\x01\x9er\x120\xb8\x05}~\x92BY\x14\xf1\xf5R\n|\xa5\xf7'\xbb\xe5,\x84\xbf\xe8\x0eH\xc3\x9b`\xc0u\xedj\x10Y\xb7\xcbu\xcf:\x8d\x93\xd6\xd0\xe3z)W*z\xd6\xc6\xb6\xd2'\xbfD\x16`]\x12\xcb\x7f[\xfc\xd0\xed\x869o\xa0\xef\xe0\xa3\xa0")

	mockUserAuthRepository := new(mysql.UserAuthRepositoryMock)

	testService := usecase.NewAuthenticationUsecase(mockUserAuthRepository, secretKey, nil)

	userAuthMockData := &model.UserAuth{
		ID:       1,
		Username: "05111740000184",
		Password: "$2a$14$JC4.1C0npGNHT8E03/O54.Clq5a/pAthGEnI01wbYFWU8p7KPqaG2",
	}

	mockUserAuthRepository.On("GetUserAuthByUsername").Return(userAuthMockData, nil)

	username := "05111740000184"
	password := "password1231"

	_, err := testService.Authenticate(username, password)

	mockUserAuthRepository.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.NotNil(t, err)

}

func TestAuthenticateUserWithUsernameDoesNotExist(t *testing.T) {

	secretKey := []byte("\xec\xbb\x81\x1fy\xff\tDi\xca\xc9\xd5\x92f{L\xadNh}fz\xe5\x04HS\x92x\x1f\xf0\xd2c,\xb0\xf2Z\xcfz\ru\x86\xfb)%\x89\xc5\x89Im\x84\xde\xeb\x15\xe6\xe5\x04A\xa5p\xeal\x97\xcb\xb7<\xb8y\xfb\xa0;V h\x0f\xc0YK\r\xa3\x8cq\x9f\x19?\xdf\n\xd8B\r \xe7s-\xd1\x1dG\x1bw\xa1\xef\x8f\xc6\xbe\x98\x90\xa7\xf4g\xc1\xcfn@\xe2\x83\x8b\xfb\xbb+\x94d\xb3\x98fD\x87\xe9\xe6m\x99\xee&_\xf9\xd1p\x99\xe7\x99}\xd9\x1b\x1fIj\x836r\xad\xff\xfd\x8dt\xcdFe\x9c\x8c\xd5S\x8a\xe2U\xad\xbd\xccw\xe6\xaf\xec\x0c\xd54?X\xf1\x15\xf1i\x01\x9er\x120\xb8\x05}~\x92BY\x14\xf1\xf5R\n|\xa5\xf7'\xbb\xe5,\x84\xbf\xe8\x0eH\xc3\x9b`\xc0u\xedj\x10Y\xb7\xcbu\xcf:\x8d\x93\xd6\xd0\xe3z)W*z\xd6\xc6\xb6\xd2'\xbfD\x16`]\x12\xcb\x7f[\xfc\xd0\xed\x869o\xa0\xef\xe0\xa3\xa0")

	mockUserAuthRepository := new(mysql.UserAuthRepositoryMock)

	testService := usecase.NewAuthenticationUsecase(mockUserAuthRepository, secretKey, nil)

	mockUserAuthRepository.On("GetUserAuthByUsername").Return(&model.UserAuth{}, fmt.Errorf("User Not Found"))

	username := "05111740000184"
	password := "password1231"

	_, err := testService.Authenticate(username, password)

	mockUserAuthRepository.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.NotNil(t, err)

}
