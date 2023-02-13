package resource

import (
	"fmt"
	"net/http"
)

var (
	DefaultEnvironment       string = "allegro.pl"
	DefaultHeaderContentType string = "application/vnd.allegro.public.v1+json"
)

var (
	apiURL = "https://api." + DefaultEnvironment + "/"
)

type resource[T any] struct {
	Response *T

	environment string
	endpoint    string
	accessToken string
}

func newResource[T any](endpoint, accessToken string) *resource[T] {
	return &resource[T]{
		environment: DefaultEnvironment,
		endpoint:    endpoint,
		accessToken: accessToken,
	}
}

func (res *resource[T]) BuildRequest() (*http.Request, error) {
	req, err := http.NewRequest("GET", apiURL+res.endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", res.accessToken))
	req.Header.Set("Content-Type", DefaultHeaderContentType)
	return req, nil
}
