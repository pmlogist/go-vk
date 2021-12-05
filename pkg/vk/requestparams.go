package vk

import (
	"net/url"
	"strings"
)

func URLValues(opts *Option) (url.Values, error) {
	values := url.Values{}

	values.Add("access_token", opts.Token)
	values.Add("v", opts.Version)

	return values, nil
}

func Concatenate(params []string, separator string) string {
	var s strings.Builder
	for i, v := range params {
		if params[i] == params[len(params)-1] {
			s.WriteString(v)
		} else {
			s.WriteString(v + separator)
		}
	}
	return s.String()
}
