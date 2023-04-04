package api

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// interface to manage APIs
type API_CLIENT interface {
	PostAPI(body []byte, url string) ([]byte, error)
	GetAPI(url string) ([]byte, error)
}

// basic struct for managing oauth based client
type api_client struct {
	BaseUrl string
	client  *http.Client
}

func GetAPIClient(baseUrl string, client *http.Client) API_CLIENT {

	return &api_client{
		BaseUrl: baseUrl,
		client:  client,
	}

}

func (cl *api_client) PostAPI(body []byte, url string) ([]byte, error) {

	res, err := cl.client.Post(cl.BaseUrl+url, "application/json", bytes.NewReader(body))

	// checking error if client is able to proceed
	if err != nil {
		return nil, err
	}

	//
	response, err := ioutil.ReadAll(res.Body)

	// checking if call executed
	if err != nil {
		return nil, err
	}

	return response, nil

}

func (cl *api_client) GetAPI(url string) ([]byte, error) {

	res, err := cl.client.Get(cl.BaseUrl + url)

	// checking error if client is able to proceed
	if err != nil {
		return nil, err
	}

	//
	response, err := ioutil.ReadAll(res.Body)

	// checking if call executed
	if err != nil {
		return nil, err
	}

	return response, nil

}
