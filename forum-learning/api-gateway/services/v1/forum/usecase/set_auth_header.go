package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/forum/domain"
)

func (usecase *forumUsecase) setAuthHeader(verifyResponse *domain.VerifyResponse, requestHeader map[string][]string) error {

	requestHeader["X-Auth-Id"] = []string{fmt.Sprint(verifyResponse.ID)}

	xAuthRoles := []string{}

	for i := 0; i < len(verifyResponse.Roles); i++ {
		xAuthRoles = append(xAuthRoles, verifyResponse.Roles[i].RoleName)
	}

	requestRoles, err := json.Marshal(xAuthRoles)

	if err != nil {
		return err
	}

	requestHeader["X-Auth-Roles"] = []string{string(requestRoles)}

	return nil
}
