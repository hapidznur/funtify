package main

import "time"

// PlayBackContext is a context
type PlayBackContext struct {
	SpotifyURI    string            `json:"uri"`
	ExternalPoint map[string]string `json:"external_url"`
	Endpoint      string            `json:"href"`
	Type          string            `json:"type"`
}

// PlayHistory currentuser
type PlayHistory struct {
	Track    Tracks          `json:"track"`
	PlayedAt time.Time       `json:"played_at"`
	Context  PlayBackContext `json:"context"`
}

// Recently played
type Recently struct {
	Items []PlayHistory `json:"items"`
}

// GetRecentlyPlayed is
func (c Client) GetRecentlyPlayed() (*Recently, error) {
	uri := c.baseURL + "me/player/recently-played"
	history := Recently{}
	err := c.get(uri, &history)
	if err != nil {
		return nil, err
	}
	return &history, nil
}
