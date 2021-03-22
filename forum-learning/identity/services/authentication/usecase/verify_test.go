package usecase_test

import (
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain/mocks/mysql"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain/model"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/usecase"
)

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

func TestVerifyTokenValid(t *testing.T) {

	secretKey := []byte("\xec\xbb\x81\x1fy\xff\tDi\xca\xc9\xd5\x92f{L\xadNh}fz\xe5\x04HS\x92x\x1f\xf0\xd2c,\xb0\xf2Z\xcfz\ru\x86\xfb)%\x89\xc5\x89Im\x84\xde\xeb\x15\xe6\xe5\x04A\xa5p\xeal\x97\xcb\xb7<\xb8y\xfb\xa0;V h\x0f\xc0YK\r\xa3\x8cq\x9f\x19?\xdf\n\xd8B\r \xe7s-\xd1\x1dG\x1bw\xa1\xef\x8f\xc6\xbe\x98\x90\xa7\xf4g\xc1\xcfn@\xe2\x83\x8b\xfb\xbb+\x94d\xb3\x98fD\x87\xe9\xe6m\x99\xee&_\xf9\xd1p\x99\xe7\x99}\xd9\x1b\x1fIj\x836r\xad\xff\xfd\x8dt\xcdFe\x9c\x8c\xd5S\x8a\xe2U\xad\xbd\xccw\xe6\xaf\xec\x0c\xd54?X\xf1\x15\xf1i\x01\x9er\x120\xb8\x05}~\x92BY\x14\xf1\xf5R\n|\xa5\xf7'\xbb\xe5,\x84\xbf\xe8\x0eH\xc3\x9b`\xc0u\xedj\x10Y\xb7\xcbu\xcf:\x8d\x93\xd6\xd0\xe3z)W*z\xd6\xc6\xb6\xd2'\xbfD\x16`]\x12\xcb\x7f[\xfc\xd0\xed\x869o\xa0\xef\xe0\xa3\xa0")

	mockUserAuthRepository := new(mysql.UserAuthRepositoryMock)

	testService := usecase.NewAuthenticationUsecase(mockUserAuthRepository, secretKey, nil)

	userAuthMockData := &model.UserAuth{
		ID:       1,
		FullName: sql.NullString{Valid: true, String: "Test Name"},
		Username: "05111740000184",
		Roles: []model.Roles{
			{1, "Client"},
		},
		EmployeeNo: sql.NullString{String: "05111740000184", Valid: true},
		IsEmployee: false,
	}

	mockUserAuthRepository.On("GetUserAuthByID").Return(userAuthMockData, nil)

	expirationDate := time.Now().Add(time.Hour * 24)

	jwtToken, err := createJWTToken(userAuthMockData.ID, expirationDate, secretKey)

	response, err := testService.Verify(jwtToken)

	mockUserAuthRepository.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.Nil(t, err)

	// Test Usecase Response
	assert.Equal(t, userAuthMockData.ID, response.ID)
	assert.Equal(t, userAuthMockData.FullName.String, response.FullName)
	assert.Equal(t, userAuthMockData.Username, response.Username)
	assert.Equal(t, userAuthMockData.Roles, response.Roles)
	assert.Equal(t, userAuthMockData.EmployeeNo, response.EmployeeNo)
	assert.Equal(t, userAuthMockData.IsEmployee, response.IsEmployee)

}

func TestVerifyTokenExpired(t *testing.T) {

	secretKey := []byte("\xec\xbb\x81\x1fy\xff\tDi\xca\xc9\xd5\x92f{L\xadNh}fz\xe5\x04HS\x92x\x1f\xf0\xd2c,\xb0\xf2Z\xcfz\ru\x86\xfb)%\x89\xc5\x89Im\x84\xde\xeb\x15\xe6\xe5\x04A\xa5p\xeal\x97\xcb\xb7<\xb8y\xfb\xa0;V h\x0f\xc0YK\r\xa3\x8cq\x9f\x19?\xdf\n\xd8B\r \xe7s-\xd1\x1dG\x1bw\xa1\xef\x8f\xc6\xbe\x98\x90\xa7\xf4g\xc1\xcfn@\xe2\x83\x8b\xfb\xbb+\x94d\xb3\x98fD\x87\xe9\xe6m\x99\xee&_\xf9\xd1p\x99\xe7\x99}\xd9\x1b\x1fIj\x836r\xad\xff\xfd\x8dt\xcdFe\x9c\x8c\xd5S\x8a\xe2U\xad\xbd\xccw\xe6\xaf\xec\x0c\xd54?X\xf1\x15\xf1i\x01\x9er\x120\xb8\x05}~\x92BY\x14\xf1\xf5R\n|\xa5\xf7'\xbb\xe5,\x84\xbf\xe8\x0eH\xc3\x9b`\xc0u\xedj\x10Y\xb7\xcbu\xcf:\x8d\x93\xd6\xd0\xe3z)W*z\xd6\xc6\xb6\xd2'\xbfD\x16`]\x12\xcb\x7f[\xfc\xd0\xed\x869o\xa0\xef\xe0\xa3\xa0")

	testService := usecase.NewAuthenticationUsecase(nil, secretKey, nil)

	expirationDate := time.Now().Add(-time.Second)

	jwtToken, err := createJWTToken(1, expirationDate, secretKey)

	_, err = testService.Verify(jwtToken)

	assert.NotNil(t, err)

}

func TestVerifyUserNotFound(t *testing.T) {

	secretKey := []byte("\xec\xbb\x81\x1fy\xff\tDi\xca\xc9\xd5\x92f{L\xadNh}fz\xe5\x04HS\x92x\x1f\xf0\xd2c,\xb0\xf2Z\xcfz\ru\x86\xfb)%\x89\xc5\x89Im\x84\xde\xeb\x15\xe6\xe5\x04A\xa5p\xeal\x97\xcb\xb7<\xb8y\xfb\xa0;V h\x0f\xc0YK\r\xa3\x8cq\x9f\x19?\xdf\n\xd8B\r \xe7s-\xd1\x1dG\x1bw\xa1\xef\x8f\xc6\xbe\x98\x90\xa7\xf4g\xc1\xcfn@\xe2\x83\x8b\xfb\xbb+\x94d\xb3\x98fD\x87\xe9\xe6m\x99\xee&_\xf9\xd1p\x99\xe7\x99}\xd9\x1b\x1fIj\x836r\xad\xff\xfd\x8dt\xcdFe\x9c\x8c\xd5S\x8a\xe2U\xad\xbd\xccw\xe6\xaf\xec\x0c\xd54?X\xf1\x15\xf1i\x01\x9er\x120\xb8\x05}~\x92BY\x14\xf1\xf5R\n|\xa5\xf7'\xbb\xe5,\x84\xbf\xe8\x0eH\xc3\x9b`\xc0u\xedj\x10Y\xb7\xcbu\xcf:\x8d\x93\xd6\xd0\xe3z)W*z\xd6\xc6\xb6\xd2'\xbfD\x16`]\x12\xcb\x7f[\xfc\xd0\xed\x869o\xa0\xef\xe0\xa3\xa0")

	mockUserAuthRepository := new(mysql.UserAuthRepositoryMock)

	testService := usecase.NewAuthenticationUsecase(mockUserAuthRepository, secretKey, nil)

	mockUserAuthRepository.On("GetUserAuthByID").Return(&model.UserAuth{}, fmt.Errorf("User Not Found"))

	expirationDate := time.Now().Add(time.Hour * 24)

	jwtToken, err := createJWTToken(1, expirationDate, secretKey)

	_, err = testService.Verify(jwtToken)

	mockUserAuthRepository.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.NotNil(t, err)

}
