package http

import (
	"fmt"
	"io"
	"net/http"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/identity/authentication/domain"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/utils"
)

type authenticationRepository struct {
	httpClient   *http.Client
	endpointsMap map[string]utils.Endpoint
}

// NewAuthenticationRepository is a constructor of authenticationRepository
// which implement AuthenticationRepository Interface
func NewAuthenticationRepository(httpClient *http.Client, endpointsMap map[string]utils.Endpoint) domain.AuthenticationRepository {
	return &authenticationRepository{
		httpClient:   httpClient,
		endpointsMap: endpointsMap,
	}
}

func (repo *authenticationRepository) Authenticate(requestBody io.Reader) (*http.Response, error) {

	endpoint, ok := repo.endpointsMap["identity"]

	if !ok {
		return nil, fmt.Errorf("Base URL of Services Not Found")
	}

	requestURL := endpoint.Host + ":" + endpoint.Port + "/auth/authenticate"

	req, err := http.NewRequest("POST", requestURL, requestBody)

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
