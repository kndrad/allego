package device

import (
	"net/http"
)

const authorizationURL = "https://allegro.pl/auth/oauth/device?client_id="

var (
	defaultContentType = "application/x-www-form-urlencoded"
)

type authorizationResponse struct {
	UserCode                string `json:"user_code"`
	DeviceCode              string `json:"device_code"`
	ExpiresIn               int    `json:"expires_in"`
	Interval                int    `json:"interval"`
	VerificationURL         string `json:"verification_uri"`
	VerificationURLComplete string `json:"verification_uri_complete"`
}

type authorization struct {
	ClientID, Secret string
	Response         *authorizationResponse
}

func newAuthorization(clientID, secret string) *authorization {
	return &authorization{
		ClientID: clientID, Secret: secret,
	}
}

func (a *authorization) BuildRequest() (*http.Request, error) {
	url := authorizationURL + a.ClientID
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", defaultContentType)
	req.SetBasicAuth(a.ClientID, a.Secret)
	return req, nil
}
