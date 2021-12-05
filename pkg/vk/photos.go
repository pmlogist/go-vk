package vk

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
)

type Photos service

type Photo struct {
	AlbumID *int64 `json:"album_id"`
}

type PhotosGetParams struct {
	OwnerId    string   `url:"owner_id"`
	AlbumID    string   `url:"album_id"`
	PhotoIds   []string `url:"photo_ids"`
	Rev        int8     `url:"rev"`
	FeedType   string   `url:"feed_type"`
	Feed       int64    `url:"feed"`
	PhotoSizes bool     `url:"photo_sizes"`
	Offset     int64    `url:"offset"`
	Count      int64    `url:"count"`
}

func (s *Photos) Get(ctx context.Context, params *PhotosGetParams) {
	method := "photos.get"
	fmt.Print(method)

	// album_id := queryparams.Concatenate(params.AlbumID, ",")
	photos_ids := Concatenate(params.PhotoIds, ",")

	query, _ := URLValues(s.option)
	query.Add("album_id", params.AlbumID)
	query.Add("photo_ids", photos_ids)
	query.Add("owner_id", params.OwnerId)

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
