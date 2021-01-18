package usecase

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func (gu *gatewayUsecase) Gateway(request *http.Request) (*http.Response, error) {

	urlRequest := request.URL.Path

	urlPrefix := strings.Split(urlRequest, "/")[1]
	urlRequestPath := strings.TrimPrefix(urlRequest, "/"+urlPrefix)

	enpoint, err := gu.enpointsRepository.FindByServicePrefix(urlPrefix)

	if err != nil {
		return nil, err
	}

	remote, err := url.Parse("http://" + enpoint.ServiceHost + enpoint.ServicePort)

	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)

	proxy.Director = func(req *http.Request) {
		req.Header = req.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = request.Param("proxyPath")
	}

	req, err := http.NewRequest(request.Method, request.RequestURI, request.Body)

	request.Header.

	req.Header.

}
