package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/identity/account_manager/domain"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/utils"
)

type accountManagerRepository struct {
	httpClient   *http.Client
	endpointsMap map[string]utils.Endpoint
}

// NewAccountManagerRepository is a constructor of accountManagerRepository
// which implement AccountManagerRepository Interface
func NewAccountManagerRepository(httpClient *http.Client, endpointsMap map[string]utils.Endpoint) domain.AccountManagerRepository {
	return &accountManagerRepository{
		httpClient:   httpClient,
		endpointsMap: endpointsMap,
	}
}

func (repo *accountManagerRepository) GetUserData(userID int) (*http.Response, error) {

	endpoint, ok := repo.endpointsMap["identity"]

	if !ok {
		return nil, fmt.Errorf("Base URL of Services Not Found")
	}

	requestURL := endpoint.Host + ":" + endpoint.Port + "/user/" + fmt.Sprint(userID)

	req, err := http.NewRequest("GET", requestURL, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := repo.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil

}

func (repo *accountManagerRepository) UpdateUserData(userID int, requestBody io.Reader, xAuthID int, xAuthRoles []string) (*http.Response, error) {

	endpoint, ok := repo.endpointsMap["identity"]

	if !ok {
		return nil, fmt.Errorf("Base URL of Services Not Found")
	}

	requestURL := endpoint.Host + ":" + endpoint.Port + "/user/update/" + fmt.Sprint(userID)

	req, err := http.NewRequest("POST", requestURL, requestBody)

	if err != nil {
		return nil, err
	}

	requestRoles, err := json.Marshal(xAuthRoles)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Id", fmt.Sprint(xAuthID))
	req.Header.Set("X-Auth-Roles", string(requestRoles))

	res, err := repo.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil

}
