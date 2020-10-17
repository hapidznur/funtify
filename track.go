package main

type Album struct {
	AlbumType        string            `json:"album_type"`
	Markets          []string          `json:"available_markets"`
	ExternalEndpoint map[string]string `json:"external_urls"`
	Endpoint         string            `json:"href"`
	ID               string            `json:"id"`
	Images           []Images          `json:"images"`
}

type Artist struct {
	ExternalEndpoint map[string]string `json:"external_urls"`
	Endpoint         string            `json:"href"`
	ID               string            `json:"id"`
	Name             string            `json:"name"`
	Artist           string            `json:"type"`
	SpotifyURI       string            `json:"uri"`
}

type Tracks struct {
	Album            Album             `json:"album"`
	Artists          []Artist          `json:"artists"`
	AvalaibleMarket  []string          `json:"available_markets"`
	DiscNumber       int               `json:"disc_number"`
	Duration         int               `json:"duration_ms"`
	Explisit         bool              `json:"explisit"`
	ExternalIDs      map[string]string `json:"external_ids"`
	ExternalEndpoint map[string]string `json:"external_url"`
	Endpoint         string            `json:"href"`
	Name             string            `json:"name"`
	Popularity       int               `json:"popularity"`
	PreviewURL       string            `json:"preview_url"`
	TrackNumber      int               `json:"track_number"`
	Type             string            `json:"type"`
	UniqueURI        string            `json:"uri"`
}
