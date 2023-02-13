package device

import (
	"allego/request"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const (
	tokenURL      = "https://allegro.pl/auth/oauth/token"
	tokenEndpoint = "?grant_type=urn:ietf:params:oauth:grant-type:device_code&device_code="
)

type deviceCode = string

type token struct {
	clientID, secret string
	code             deviceCode

	response *tokenResponse
}

func newToken(clientID, secret string, code deviceCode) *token {
	return &token{
		clientID: clientID, secret: secret,
		code: code,
	}
}

func (t *token) BuildRequest() (*http.Request, error) {
	url := tokenURL + tokenEndpoint + t.code
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(t.clientID, t.secret)
	return req, nil
}

func (t *token) StartPolling(interval int) (*tokenResponse, error) {
	for {
		resp, err := request.Send(http.DefaultClient, t)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if err = json.NewDecoder(resp.Body).Decode(&t.response); err != nil {
			log.Fatal(err)
		}

		switch resp.StatusCode {
		case http.StatusBadRequest:
			switch t.response.Error {
			case "slow_down":
				interval += interval
			case "access_denied":
				break
			}
		case http.StatusOK:
			return t.response, nil
		}
		time.Sleep(time.Second * time.Duration(interval))
	}
}

type tokenResponse struct {
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"refresh_token"`
	TokenType    string  `json:"token_type"`
	ExpiresIn    float64 `json:"expires_in"`
	Scope        string  `json:"scope"`
	AllegroAPI   bool    `json:"allegro_api"`
	JTI          string  `json:"jti"`

	Error string `json:"error,omitempty"`
}
