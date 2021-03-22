package usecase

import (
	"fmt"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain"
)

func (usecase *authenticationUsecase) Verify(token string) (*domain.VerifyUsecaseResponse, error) {

	jwtTokenObject, err := usecase.VerifyToken(token, usecase.secretKey)

	if err != nil {
		return nil, err
	}

	claims, ok := jwtTokenObject.Claims.(jwt.MapClaims)

	if !ok || !jwtTokenObject.Valid {
		return nil, fmt.Errorf("Token Invalid")
	}

	userID, ok := claims["id"].(string)

	if !ok {
		return nil, fmt.Errorf("Token Metadata Not Found")
	}

	userIDInt, err := strconv.Atoi(userID)

	if err != nil {
		return nil, fmt.Errorf("ID Format Invalid")
	}

	userAuth, err := usecase.userAuthRepository.GetUserAuthByID(userIDInt)

	if err != nil {
		return nil, err
	}

	response := &domain.VerifyUsecaseResponse{
		ID:         userAuth.ID,
		FullName:   userAuth.FullName.String,
		Username:   userAuth.Username,
		Roles:      userAuth.Roles,
		EmployeeNo: userAuth.EmployeeNo.String,
		IsEmployee: userAuth.IsEmployee,
	}

	return response, nil

}
