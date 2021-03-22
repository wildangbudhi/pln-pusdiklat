package usecase

import (
	"fmt"
	"strings"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/identity/account_manager/domain"
)

func (usecase *accountManagerUsecase) Authenticate(token string) (*domain.VerifyResponse, error) {

	tokenArray := strings.Fields(token)

	if len(tokenArray) != 2 {
		return nil, fmt.Errorf("Token Format Not Valid")
	}

	tokenType := tokenArray[0]
	tokenString := tokenArray[1]

	var verifyRespones *domain.VerifyResponse
	var err error

	if tokenType == "native" {
		verifyRespones, err = usecase.authenticationRepository.Verify(tokenString)
	} else if tokenType == "sso" {
		verifyRespones, err = usecase.authenticationRepository.VerifySSOPLN(tokenString)
	}

	if err != nil {
		return nil, err
	}

	return verifyRespones, nil

}
