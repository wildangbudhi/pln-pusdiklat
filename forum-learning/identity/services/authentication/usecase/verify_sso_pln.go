package usecase

import (
	"database/sql"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain/model"
)

func (usecase *authenticationUsecase) VerifySSOPLN(token string) (*domain.VerifyUsecaseResponse, error) {

	jwtTokenObject, err := usecase.VerifyToken(token, usecase.apiSecretKey)

	if err != nil {
		return nil, err
	}

	claims, ok := jwtTokenObject.Claims.(jwt.MapClaims)

	if !ok || !jwtTokenObject.Valid {
		return nil, fmt.Errorf("Token Invalid")
	}

	employeeNo, ok := claims["employee_no"].(string)

	if !ok {
		return nil, fmt.Errorf("Token Metadata Not Invalid")
	}

	fullName, ok := claims["full_name"].(string)

	if !ok {
		return nil, fmt.Errorf("Token Metadata Not Invalid")
	}

	username, ok := claims["username"].(string)

	if !ok {
		return nil, fmt.Errorf("Token Metadata Not Invalid")
	}

	userAuth, err := usecase.userAuthRepository.GetUserAuthByUsername(username)

	if err != nil {

		if err.Error() != "User Not Found" {
			return nil, err
		} else {

			userAuth = &model.UserAuth{
				FullName: sql.NullString{String: fullName, Valid: true},
				Username: username,
				Roles: []model.Roles{
					{ID: 1, RoleName: "Client"},
				},
				EmployeeNo: sql.NullString{String: employeeNo, Valid: true},
				IsEmployee: true,
			}

			userID, err := usecase.userAuthRepository.InsertUserAuth(userAuth)

			if err != nil {
				return nil, err
			}

			userAuth.ID = int(userID)

		}
	}

	result := &domain.VerifyUsecaseResponse{
		ID:         userAuth.ID,
		FullName:   userAuth.FullName.String,
		Username:   userAuth.Username,
		Roles:      userAuth.Roles,
		EmployeeNo: userAuth.EmployeeNo.String,
		IsEmployee: userAuth.IsEmployee,
	}

	return result, nil

}
