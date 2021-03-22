package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/identity/account_manager/domain"
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

func (repo *authenticationRepository) Verify(token string) (*domain.VerifyResponse, error) {

	endpoint, ok := repo.endpointsMap["identity"]

	if !ok {
		return nil, fmt.Errorf("Base URL of Services Not Found")
	}

	requestURL := endpoint.Host + ":" + endpoint.Port + "/auth/verify"

	reqData := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}

	jsonByte, err := json.Marshal(reqData)

	if err != nil {
		return nil, err
	}

	bodyObj := bytes.NewBuffer(jsonByte)

	req, err := http.NewRequest("POST", requestURL, bodyObj)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := repo.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	responseObj := &domain.VerifyResponse{}

	err = json.NewDecoder(res.Body).Decode(responseObj)

	if err != nil {

		errorResponse := &domain.HTTPErrorResponse{}
		err = json.NewDecoder(res.Body).Decode(errorResponse)

		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf(errorResponse.Error)

	}

	return responseObj, nil

}

func (repo *authenticationRepository) VerifySSOPLN(token string) (*domain.VerifyResponse, error) {

	endpoint, ok := repo.endpointsMap["identity"]

	if !ok {
		return nil, fmt.Errorf("Base URL of Services Not Found")
	}

	requestURL := endpoint.Host + ":" + endpoint.Port + "/auth/verify/sso/pln"

	reqData := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}

	jsonByte, err := json.Marshal(reqData)

	if err != nil {
		return nil, err
	}

	bodyObj := bytes.NewBuffer(jsonByte)

	req, err := http.NewRequest("POST", requestURL, bodyObj)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := repo.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	responseObj := &domain.VerifyResponse{}

	err = json.NewDecoder(res.Body).Decode(responseObj)

	if err != nil {

		errorResponse := &domain.HTTPErrorResponse{}
		err = json.NewDecoder(res.Body).Decode(errorResponse)

		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf(errorResponse.Error)

	}

	return responseObj, nil

}

func (repo *authenticationRepository) EndpointAuthorize(data *domain.EndpointAuthorizeParameter) (*domain.EndpointAuthorizeResponse, error) {

	endpoint, ok := repo.endpointsMap["identity"]

	if !ok {
		return nil, fmt.Errorf("Base URL of Services Not Found")
	}

	requestURL := endpoint.Host + ":" + endpoint.Port + "/auth/endpoint/authorize"

	jsonByte, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	bodyObj := bytes.NewBuffer(jsonByte)

	req, err := http.NewRequest("POST", requestURL, bodyObj)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := repo.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	responseObj := &domain.EndpointAuthorizeResponse{}

	err = json.NewDecoder(res.Body).Decode(responseObj)

	if err != nil {
		errorResponse := &domain.HTTPErrorResponse{}
		err = json.NewDecoder(res.Body).Decode(errorResponse)

		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf(errorResponse.Error)
	}

	return responseObj, nil

}
