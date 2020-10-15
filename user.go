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

// User model
type User struct {
	// DisplayedName
	DisplayName  string            `json:"display_name"`
	ExternalURLs map[string]string `json:"external_urls`
	Followers    string            `json:"followers"`
	Endpoint     string            `json:href`
	ID           string            `json:"id"`
	Images       []string          `json:"images"`
}

// PrivateUser CurrentUser
type PrivateUser struct {
	User
	// The country ISO 3166-1 alpha code
	Country string `json:"country"`
	Email   string `json:"email"`
	// Product subcription level "open"
	Product   string `json:"product"`
	Birthdate string `json:"birthdate"`
}

// CurrentUser to get
func (client *Client) CurrentUser() (*PrivateUser, error) {
	spotifyURL := client.baseURL + "me"
	var result PrivateUser

	req, err := http.NewRequest("GET", spotifyURL, nil)
	if client.AcceptLanguage != "" {
		req.Header.Set("Accept-Language", client.AcceptLanguage)
	}

	if err != nil {
		return nil, err
	}

	resp, err := client.http.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Print(resp.StatusCode)
	}

	json.NewDecoder(resp.Body).Decode(&result)
	return &result, nil
}
