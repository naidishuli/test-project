package interfaces

import "net/http"

type HTTPClient interface {
	Do(r *http.Request) (*http.Response, error)
}
