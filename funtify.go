package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BasePage struct {
	Endpoint  string `json:"href"`
	Limit     string `json:"limit"`
	Nextpoint string `json:"next"`
	Offset    uint   `json:"offset"`
	Prevpoint string `json:"previous"`
	Count     uint   `json:"total"`
}

type Images struct {
	Height   int    `json:"height"`
	Endpoint string `json:"href"`
	Width    int    `json:"Width"`
}

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
