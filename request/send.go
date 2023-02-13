package request

import (
	"net/http"
)

// RequestBuilder interface describes BuildRequest() method which builds http.Request.
type RequestBuilder interface {
	BuildRequest() (*http.Request, error)
}

// Send takes pointer http.Client and RequestBuilder interface.
func Send(client *http.Client, rb RequestBuilder) (*http.Response, error) {
	req, err := rb.BuildRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
