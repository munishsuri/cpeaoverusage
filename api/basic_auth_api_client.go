package api

import (
	"bytes"
	b64 "encoding/base64"
	"io"
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// interface to manage APIs
type BASIC_API_CLIENT interface {
	PostAPI(body []byte, url string) ([]byte, error)
}

// basic struct for managing oauth based client
type basic_api_client struct {
	BaseUrl  string
	User     string
	Password string
	Client   HTTPClient
}

func GetBasicAPIClient(baseUrl, username, password string, client *http.Client) BASIC_API_CLIENT {

	return &basic_api_client{
		BaseUrl:  baseUrl,
		User:     username,
		Password: password,
		Client:   client,
	}

}

func (cl *basic_api_client) PostAPI(body []byte, url string) ([]byte, error) {

	posturl := cl.BaseUrl + url

	// creating auth header
	authHeader := b64.StdEncoding.EncodeToString([]byte(cl.User + ":" + cl.Password))

	// Create a HTTP post request
	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	r.Header.Add("Content-Type", "application/json;charset=utf-8")
	r.Header.Add("Accept", "application/json;charset=utf-8")
	r.Header.Add("Authorization", "Basic "+authHeader)

	res, err := cl.Client.Do(r)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
