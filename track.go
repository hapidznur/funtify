package main

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

type TracksPage struct {
	BasePage
	Tracks []Tracks `json:"items"`
}
