package funtify

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Images struct {
	Height   int    `json:"height"`
	Endpoint string `json:"href"`
	Width    int    `json:"Width"`
}

// URI identifies an artist, album, track, or category.  For example,
// spotify:track:6rqhFgbbKwnb9MLmUQDhG6
type URI string

// ID is a base-62 identifier for an artist, track, album, etc.
// It can be found at the end of a spotify.URI.
type ID string

const (
	// DateLayout can be used with time.Parse to create time.Time values
	// from Spotify date strings.  For example, PrivateUser.Birthdate
	// uses this format.
	DateLayout = "2006-01-02"
	// TimestampLayout can be used with time.Parse to create time.Time
	// values from SpotifyTimestamp strings.  It is an ISO 8601 UTC timestamp
	// with a zero offset.  For example, PlaylistTrack's AddedAt field uses
	// this format.
	TimestampLayout = "2006-01-02T15:04:05Z"
)

// Get Htpp function to spotify
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
