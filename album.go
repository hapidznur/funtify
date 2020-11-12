package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Album struct {
	Name                 string            `json:"name"`
	Artist               []Artist          `json:"artist"`
	AlbumGroup           string            `json:"album_group"`
	AlbumType            string            `json:"album_type"`
	ID                   ID                `json:"id"`
	URI                  URI               `json:"uri"`
	Markets              []string          `json:"available_markets"`
	ExternalEndpoint     map[string]string `json:"external_urls"`
	Endpoint             string            `json:"href"`
	Images               []Images          `json:"images"`
	ReleaseDate          string            `json:"release_date"`
	ReleaseDatePrecision string            `json:"release_date_precision"`
}

func (a *Album) ReleaseDateTime() time.Time {
	if a.ReleaseDatePrecision == "day" {
		result, _ := time.Parse(DateLayout, a.ReleaseDate)
		return result
	}

	if a.ReleaseDatePrecision == "month" {
		ym := strings.Split(a.ReleaseDate, "-")
		year, _ := strconv.Atoi(ym[0])
		month, _ := strconv.Atoi(ym[1])
		return time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	}
	year, _ := strconv.Atoi(a.ReleaseDate)
	return time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
}

type Copyright struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

type FullAlbum struct {
	Album
	Copyright   []Copyright       `json:"copyrights"`
	Genres      []string          `json:"genres"`
	Popularity  int               `json:"popularity"`
	Tracks      TracksPage        `json:"tracks"`
	ExternalIDs map[string]string `json:"external_ids"`
}

func (c *Client) GetAlbum(id ID) (*FullAlbum, error) {
	var fa FullAlbum
	spotifyURL := fmt.Sprintf("%salbums/%s", c.baseURL, id)

	err := c.get(spotifyURL, &fa)
	if err != nil {
		return nil, err
	}

	return &fa, nil
}
