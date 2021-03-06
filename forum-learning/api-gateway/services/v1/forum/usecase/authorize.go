package usecase

import "github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/forum/domain"

func (usecase *forumUsecase) Authorize(userID int, method string, url string) (*domain.EndpointAuthorizeResponse, error) {

	payload := domain.EndpointAuthorizeParameter{
		UserID: userID,
		Method: method,
		URL:    url,
	}

	endpointAuthorizeRespones, err := usecase.authenticationRepository.EndpointAuthorize(&payload)

	if err != nil {
		return nil, err
	}

	return endpointAuthorizeRespones, nil

}
