package http

import (
	"fmt"
	"io"
	"net/http"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/services/v1/forum/domain"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/api-gateway/utils"
)

type forumRepository struct {
	httpClient   *http.Client
	endpointsMap map[string]utils.Endpoint
}

// NewForumRepository is a constructor of forumRepository
// which implement ForumRepository Interface
func NewForumRepository(httpClient *http.Client, endpointsMap map[string]utils.Endpoint) domain.ForumRepository {
	return &forumRepository{
		httpClient:   httpClient,
		endpointsMap: endpointsMap,
	}
}

func (repo *forumRepository) sendRequest(requestURL string, requestMethod string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error) {

	if urlQuery != "" {
		requestURL += "?" + urlQuery
	}

	req, err := http.NewRequest(requestMethod, requestURL, requestBody)

	if err != nil {
		return nil, err
	}

	for key, values := range requestHeader {
		for i := 0; i < len(values); i++ {
			req.Header.Add(key, values[i])
		}
	}

	res, err := repo.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil

}

func (repo *forumRepository) FetchCategory(requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error) {

	endpoint, ok := repo.endpointsMap["forum"]

	if !ok {
		return nil, fmt.Errorf("Base URL of Services Not Found")
	}

	requestURL := endpoint.Host + ":" + endpoint.Port + "/category"

	res, err := repo.sendRequest(requestURL, "GET", requestBody, urlQuery, requestHeader)

	if err != nil {
		return nil, err
	}

	return res, nil

}

func (repo *forumRepository) CreateForum(requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error) {

	endpoint, ok := repo.endpointsMap["forum"]

	if !ok {
		return nil, fmt.Errorf("Base URL of Services Not Found")
	}

	requestURL := endpoint.Host + ":" + endpoint.Port + "/create"

	res, err := repo.sendRequest(requestURL, "POST", requestBody, urlQuery, requestHeader)

	if err != nil {
		return nil, err
	}

	return res, nil

}

func (repo *forumRepository) UpdateForum(forumID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error) {

	endpoint, ok := repo.endpointsMap["forum"]

	if !ok {
		return nil, fmt.Errorf("Base URL of Services Not Found")
	}

	requestURL := endpoint.Host + ":" + endpoint.Port + "/update/" + forumID

	res, err := repo.sendRequest(requestURL, "POST", requestBody, urlQuery, requestHeader)

	if err != nil {
		return nil, err
	}

	return res, nil

}

func (repo *forumRepository) CloseForum(forumID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error) {

	endpoint, ok := repo.endpointsMap["forum"]

	if !ok {
		return nil, fmt.Errorf("Base URL of Services Not Found")
	}

	requestURL := endpoint.Host + ":" + endpoint.Port + "/close/" + forumID

	res, err := repo.sendRequest(requestURL, "GET", requestBody, urlQuery, requestHeader)

	if err != nil {
		return nil, err
	}

	return res, nil

}

func (repo *forumRepository) DeleteForum(forumID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error) {

	endpoint, ok := repo.endpointsMap["forum"]

	if !ok {
		return nil, fmt.Errorf("Base URL of Services Not Found")
	}

	requestURL := endpoint.Host + ":" + endpoint.Port + "/delete/" + forumID

	res, err := repo.sendRequest(requestURL, "GET", requestBody, urlQuery, requestHeader)

	if err != nil {
		return nil, err
	}

	return res, nil

}

func (repo *forumRepository) FetchWithPagination(requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error) {

	endpoint, ok := repo.endpointsMap["forum"]

	if !ok {
		return nil, fmt.Errorf("Base URL of Services Not Found")
	}

	requestURL := endpoint.Host + ":" + endpoint.Port + "/fetch"

	res, err := repo.sendRequest(requestURL, "GET", requestBody, urlQuery, requestHeader)

	if err != nil {
		return nil, err
	}

	return res, nil

}

func (repo *forumRepository) GetForum(forumID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error) {

	endpoint, ok := repo.endpointsMap["forum"]

	if !ok {
		return nil, fmt.Errorf("Base URL of Services Not Found")
	}

	requestURL := endpoint.Host + ":" + endpoint.Port + "/get/" + forumID

	res, err := repo.sendRequest(requestURL, "GET", requestBody, urlQuery, requestHeader)

	if err != nil {
		return nil, err
	}

	return res, nil

}

func (repo *forumRepository) FetchWithPaginationByAuthorID(authorID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error) {

	endpoint, ok := repo.endpointsMap["forum"]

	if !ok {
		return nil, fmt.Errorf("Base URL of Services Not Found")
	}

	requestURL := endpoint.Host + ":" + endpoint.Port + "/author/" + authorID

	res, err := repo.sendRequest(requestURL, "GET", requestBody, urlQuery, requestHeader)

	if err != nil {
		return nil, err
	}

	return res, nil

}

func (repo *forumRepository) SearchForum(requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error) {

	endpoint, ok := repo.endpointsMap["forum"]

	if !ok {
		return nil, fmt.Errorf("Base URL of Services Not Found")
	}

	requestURL := endpoint.Host + ":" + endpoint.Port + "/search"

	res, err := repo.sendRequest(requestURL, "GET", requestBody, urlQuery, requestHeader)

	if err != nil {
		return nil, err
	}

	return res, nil

}

func (repo *forumRepository) ReactForum(forumID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error) {

	endpoint, ok := repo.endpointsMap["forum"]

	if !ok {
		return nil, fmt.Errorf("Base URL of Services Not Found")
	}

	requestURL := endpoint.Host + ":" + endpoint.Port + "/react/" + forumID

	res, err := repo.sendRequest(requestURL, "GET", requestBody, urlQuery, requestHeader)

	if err != nil {
		return nil, err
	}

	return res, nil

}

func (repo *forumRepository) ReplyForum(forumID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error) {

	endpoint, ok := repo.endpointsMap["forum"]

	if !ok {
		return nil, fmt.Errorf("Base URL of Services Not Found")
	}

	requestURL := endpoint.Host + ":" + endpoint.Port + "/reply/create/" + forumID

	res, err := repo.sendRequest(requestURL, "POST", requestBody, urlQuery, requestHeader)

	if err != nil {
		return nil, err
	}

	return res, nil

}

func (repo *forumRepository) UpdateForumReplies(forumReplyID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error) {

	endpoint, ok := repo.endpointsMap["forum"]

	if !ok {
		return nil, fmt.Errorf("Base URL of Services Not Found")
	}

	requestURL := endpoint.Host + ":" + endpoint.Port + "/reply/update/" + forumReplyID

	res, err := repo.sendRequest(requestURL, "POST", requestBody, urlQuery, requestHeader)

	if err != nil {
		return nil, err
	}

	return res, nil

}

func (repo *forumRepository) DeleteForumReplies(forumReplyID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error) {

	endpoint, ok := repo.endpointsMap["forum"]

	if !ok {
		return nil, fmt.Errorf("Base URL of Services Not Found")
	}

	requestURL := endpoint.Host + ":" + endpoint.Port + "/reply/delete/" + forumReplyID

	res, err := repo.sendRequest(requestURL, "GET", requestBody, urlQuery, requestHeader)

	if err != nil {
		return nil, err
	}

	return res, nil

}

func (repo *forumRepository) ReactForumReplies(forumReplyID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error) {

	endpoint, ok := repo.endpointsMap["forum"]

	if !ok {
		return nil, fmt.Errorf("Base URL of Services Not Found")
	}

	requestURL := endpoint.Host + ":" + endpoint.Port + "/reply/react/" + forumReplyID

	res, err := repo.sendRequest(requestURL, "GET", requestBody, urlQuery, requestHeader)

	if err != nil {
		return nil, err
	}

	return res, nil

}

func (repo *forumRepository) FetchReplyByForumIDWithPagination(forumID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error) {

	endpoint, ok := repo.endpointsMap["forum"]

	if !ok {
		return nil, fmt.Errorf("Base URL of Services Not Found")
	}

	requestURL := endpoint.Host + ":" + endpoint.Port + "/get/" + forumID + "/replies"

	res, err := repo.sendRequest(requestURL, "GET", requestBody, urlQuery, requestHeader)

	if err != nil {
		return nil, err
	}

	return res, nil

}
