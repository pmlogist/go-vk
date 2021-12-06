package vk

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	client *http.Client
	common service

	BaseURL *url.URL
	API     API
}

type UsersGetResponse struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (c *Client) NewRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req, nil

}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		log.Fatal(resp)
	}

	return resp, nil
}
