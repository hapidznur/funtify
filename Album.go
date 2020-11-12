package main

type Album struct {
	AlbumType        string            `json:"album_type"`
	Markets          []string          `json:"available_markets"`
	ExternalEndpoint map[string]string `json:"external_urls"`
	Endpoint         string            `json:"href"`
	ID               string            `json:"id"`
	Images           []Images          `json:"images"`
}

type Copyright struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

type FullAlbum struct {
	BasePage
	Copyright   []Copyright       `json:"copyrights"`
	Genres      []string          `json:"genres"`
	Popularity  int               `json:"popularity"`
	Tracks      TracksPage        `json:"tracks"`
	ExternalIDs map[string]string `json:"external_ids"`
}

type AlbumPage struct {
	BasePage
	Albums []Album `json:"items"`
}
