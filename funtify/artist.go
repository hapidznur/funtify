package funtify

import "fmt"

type Artist struct {
	ExternalEndpoint map[string]string `json:"external_urls"`
	Endpoint         string            `json:"href"`
	ID               ID                `json:"id"`
	Name             string            `json:"name"`
	SpotifyURI       URI               `json:"uri"`
}

type FullArtist struct {
	Artist
	Popularity int       `json:"popularity"`
	Genres     []string  `json:"genres"`
	Followers  Followers `json:"followers"`
	Images     []Images  `json:"images"`
}

func (c *Client) GetArtist(id ID) (*FullArtist, error) {
	spotifyURL := fmt.Sprintf("%sartists/%s", c.baseURL, id)

	var artist FullArtist
	err := c.get(spotifyURL, &artist)

	if err != nil {
		return nil, err
	}
	return &artist, nil
}
