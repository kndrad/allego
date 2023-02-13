package device

import (
	"allego/browser"
	"allego/request"
	"encoding/json"
	"log"
	"net/http"
)

func Token(client *http.Client, clientID, secret string) (*tokenResponse, error) {
	authorization, err := authorize(client, clientID, secret)
	if err != nil {
		return nil, err
	}
	token := newToken(clientID, secret, authorization.Response.DeviceCode)
	go browser.OpenURL(authorization.Response.VerificationURLComplete)
	resp, err := token.StartPolling(authorization.Response.Interval)
	if err != nil {
		log.Fatal(err)
	}
	return resp, nil
}

func authorize(client *http.Client, clientID, secret string) (*authorization, error) {
	authorization := newAuthorization(clientID, secret)
	resp, err := request.Send(client, authorization)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err = json.NewDecoder(resp.Body).Decode(&authorization.Response); err != nil {
		return nil, err
	}
	return authorization, nil
}
