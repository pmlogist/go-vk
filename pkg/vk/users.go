package vk

import (
	"context"
	"encoding/json"
	"io"
	"log"
)

type Users service

// MethodError represents a VK API method call error.
type Error struct {
	Code          int64  `json:"error_code"`
	Message       string `json:"error_msg"`
	RequestParams []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"request_params"`
}
type UsersGetParams struct {
	UserIds  []string `url:"user_ids"`
	Fields   []string `url:"fields"`
	NameCase string   `url:"name_case"`
}

func (s *Users) Get(ctx context.Context, params *UsersGetParams) ([]UserFull, error) {
	method := "users.get"
	var body struct {
		Response []UserFull `json:"response"`
		Error    *Error     `json:"error"`
	}

	fields := Concatenate(params.Fields, ",")
	userIds := Concatenate(params.UserIds, ",")

	query, _ := URLValues(s.option)
	query.Add("fields", fields)
	query.Add("user_ids", userIds)

	url := (s.client.BaseURL.String() + "/" + method + "?" + query.Encode())

	req, _ := s.client.NewRequest("POST", url, nil)
	resp, err := s.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	rawBody, _ := io.ReadAll(resp.Body)

	if err = json.Unmarshal(rawBody, &body); err != nil {
		log.Fatal(err)
	}

	return body.Response, err

}
