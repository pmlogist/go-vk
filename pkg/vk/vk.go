package vk

import (
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
