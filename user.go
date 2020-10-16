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
	ExternalURLs map[string]string `json:"external_urls"`
	Followers    Followers         `json:"followers"`
	Endpoint     string            `json:"href"`
	ID           string            `json:"id"`
	Images       []string          `json:"images"`
}

// Follower data
type Followers struct {
	Endpoint string `json:"href"`
	Total    uint   `json:"total"`
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
	err := client.get(spotifyURL, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Get function to spotify
func (c *Client) get(url string, result interface{}) error {

	req, err := http.NewRequest("GET", url, nil)
	if c.AcceptLanguage != "" {
		req.Header.Set("Accept-Language", c.AcceptLanguage)
	}

	if err != nil {
		return err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Print(resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return err
	}

	return nil
}
