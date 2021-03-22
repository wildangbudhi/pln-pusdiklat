package usecase_test

import (
	"database/sql"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain/mocks/mysql"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain/model"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/usecase"
)

func createSAMLJWTToken(employeeNo string, fullName string, username string, secretKey []byte) (string, error) {

	aksesTokenClaims := jwt.MapClaims{
		"employee_no": employeeNo,
		"full_name":   fullName,
		"username":    username,
	}

	aksesToken := jwt.NewWithClaims(jwt.SigningMethodHS512, aksesTokenClaims)

	token, err := aksesToken.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return token, nil

}

func TestVerifySSOPLNTokenValid(t *testing.T) {

	secretAPIKey := []byte("\xec\xbb\x81\x1fy\xff\tDi\xca\xc9\xd5\x92f{L\xadNh}fz\xe5\x04HS\x92x\x1f\xf0\xd2c,\xb0\xf2Z\xcfz\ru\x86\xfb)%\x89\xc5\x89Im\x84\xde\xeb\x15\xe6\xe5\x04A\xa5p\xeal\x97\xcb\xb7<\xb8y\xfb\xa0;V h\x0f\xc0YK\r\xa3\x8cq\x9f\x19?\xdf\n\xd8B\r \xe7s-\xd1\x1dG\x1bw\xa1\xef\x8f\xc6\xbe\x98\x90\xa7\xf4g\xc1\xcfn@\xe2\x83\x8b\xfb\xbb+\x94d\xb3\x98fD\x87\xe9\xe6m\x99\xee&_\xf9\xd1p\x99\xe7\x99}\xd9\x1b\x1fIj\x836r\xad\xff\xfd\x8dt\xcdFe\x9c\x8c\xd5S\x8a\xe2U\xad\xbd\xccw\xe6\xaf\xec\x0c\xd54?X\xf1\x15\xf1i\x01\x9er\x120\xb8\x05}~\x92BY\x14\xf1\xf5R\n|\xa5\xf7'\xbb\xe5,\x84\xbf\xe8\x0eH\xc3\x9b`\xc0u\xedj\x10Y\xb7\xcbu\xcf:\x8d\x93\xd6\xd0\xe3z)W*z\xd6\xc6\xb6\xd2'\xbfD\x16`]\x12\xcb\x7f[\xfc\xd0\xed\x869o\xa0\xef\xe0\xa3\xa0")

	mockUserAuthRepository := new(mysql.UserAuthRepositoryMock)

	testService := usecase.NewAuthenticationUsecase(mockUserAuthRepository, nil, secretAPIKey)

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

	mockUserAuthRepository.On("GetUserAuthByUsername").Return(userAuthMockData, nil)

	jwtToken, err := createSAMLJWTToken(userAuthMockData.EmployeeNo.String, userAuthMockData.FullName.String, userAuthMockData.Username, secretAPIKey)

	response, err := testService.VerifySSOPLN(jwtToken)

	mockUserAuthRepository.AssertExpectations(t)

	// Test Usecase Error is Nil
	assert.Nil(t, err)

	// Test Usecase Response
	assert.Equal(t, userAuthMockData.ID, response.ID)
	assert.Equal(t, userAuthMockData.FullName.String, response.FullName)
	assert.Equal(t, userAuthMockData.Username, response.Username)
	assert.Equal(t, userAuthMockData.Roles, response.Roles)
	assert.Equal(t, userAuthMockData.EmployeeNo.String, response.EmployeeNo)
	assert.Equal(t, userAuthMockData.IsEmployee, response.IsEmployee)

}
