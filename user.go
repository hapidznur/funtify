package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const baseAddress = "https://api.spotify.com/v1/"

// Client struct to
type Client struct {
	http    *http.Client
	baseURL string

	AutoRetry      bool
	AcceptLanguage string
}

// CurrentUser to get
func (client *Client) CurrentUser(userID string) {
	spotifyURL := baseAddress + "user/" + string(userID)
	fmt.Print(spotifyURL)
	var user string

	req, err := http.NewRequest("GET", spotifyURL, nil)
	if client.AcceptLanguage != "" {
		req.Header.Set("Accept-Language", client.AcceptLanguage)
	}
	if err != nil {
		fmt.Print(err)
	}
	resp, err := client.http.Do(req)
	if err != nil {
		fmt.Print(err)
	}

	defer resp.Body.Close()
	result := json.NewDecoder(resp.Body).Decode(user)
	fmt.Print(result)
}
