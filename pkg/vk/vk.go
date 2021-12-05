package vk

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://api.vk.com/method"
	defaultVersion = "5.131"
)

type Option struct {
	Token   string `url:"access_token"`
	Version string `url:"v"`
}

type service struct {
	client *Client
	option *Option
}

type API struct {
	Photos *Photos
	Users  *Users
}

type Objects struct {
}

type Client struct {
	client *http.Client
	common service

	BaseURL *url.URL
	API     API
}

func (c *Client) NewRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	return req, nil

}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		log.Fatal(resp)
	}

	return resp, nil
}

func New(opts *Option) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		client:  &http.Client{},
		BaseURL: baseURL,
	}

	c.common.client = c
	c.common.option = opts
	if opts.Version == "" {
		c.common.option.Version = defaultVersion
	}

	c.API.Photos = (*Photos)(&c.common)
	c.API.Users = (*Users)(&c.common)

	return c
}
