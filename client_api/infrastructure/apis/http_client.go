package apis

import (
	"net/http"
)

type httpClient struct {
	*http.Client
}

func NewHttpClient() *httpClient {
	return &httpClient{
		&http.Client{},
	}
}

func (h *httpClient) Do(r *http.Request) (*http.Response, error) {
	return h.Client.Do(r)
}
