package vk

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
)

type Users service

type UsersGetParams struct {
	UserIds  []string `url:"user_ids"`
	Fields   []string `url:"fields"`
	NameCase string   `url:"name_case"`
}

func (s *Users) Get(ctx context.Context, params *UsersGetParams) {
	method := "users.get"

	fields := Concatenate(params.Fields, ",")
	userIds := Concatenate(params.UserIds, ",")

	query, _ := URLValues(s.option)
	query.Add("fields", fields)
	query.Add("user_ids", userIds)

	url := (s.client.BaseURL.String() + "/" + method + "?" + query.Encode())

	req, _ := s.client.NewRequest("GET", url, nil)

	resp, err := s.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	sb := string(body)
	fmt.Println(sb)
}
