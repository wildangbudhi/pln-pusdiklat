package usecase

import (
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/account_manager/domain"
)

func (usecase *accountManagerUsecase) GetUserData(userID int) (*domain.GetUserDataUsecaseResponse, error) {

	userAuthData, err := usecase.userAuthRepository.GetUserAuthByID(userID)

	if err != nil {
		return nil, err
	}

	response := &domain.GetUserDataUsecaseResponse{
		ID:         userAuthData.ID,
		FullName:   userAuthData.FullName.String,
		AvatarFile: userAuthData.AvatarFile.String,
		Email:      userAuthData.Email,
		Username:   userAuthData.Username,
	}

	return response, nil

}
