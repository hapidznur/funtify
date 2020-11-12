package main

type Artist struct {
	ExternalEndpoint map[string]string `json:"external_urls"`
	Endpoint         string            `json:"href"`
	ID               string            `json:"id"`
	Name             string            `json:"name"`
	Artist           string            `json:"type"`
	SpotifyURI       string            `json:"uri"`
}
