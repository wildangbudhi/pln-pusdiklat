package domain

import "net/http"

type GatewayUsecase interface {
	Gateway(request *http.Request) (*http.Response, error)
}
