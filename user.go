package main

import (
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

type ItemSavedTracks struct {
	Added string   `json:"added_at"`
	track []Tracks `json:"tracks"`
}

type SavedTracks struct {
	BasePage
	ItemSavedTracks
}

// CurrentUser to get
func (c *Client) CurrentUser() (*PrivateUser, error) {
	spotifyURL := c.baseURL + "me"
	var result PrivateUser
	err := c.get(spotifyURL, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UserProfile is
func (c *Client) UserProfile(userID string) (*User, error) {
	var user User
	spotifyURL := c.baseURL + "users/" + userID
	err := c.get(spotifyURL, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (c *Client) CurrentUserTracks() error {
	spotifyURL := c.baseURL + "me/tracks"
	var savedTracks SavedTracks
	err := c.get(spotifyURL, savedTracks)
	if err != nil {
		return err
	}
	return nil
}
