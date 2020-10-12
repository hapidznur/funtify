package main

import (
	"context"
	"crypto/tls"
	"errors"
	"net/http"

	"golang.org/x/oauth2"
)

const (
	AuthURL  = "https://accounts.spotify.com/authorize"
	TokenURL = "https://accounts.spotify.com/api/token"
)

const (
	ScopeUserReadPrivate = "user-read-private"
)

type Authenticator struct {
	config  *oauth2.Config
	context context.Context
}

func Authorize(redirectURL string, scopes ...string) Authenticator {
	cfg := &oauth2.Config{
		ClientID:     readEnvVariable("spotify_id"),
		ClientSecret: readEnvVariable("spotify_secret"),
		RedirectURL:  redirectURL,
		Scopes:       scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  AuthURL,
			TokenURL: TokenURL,
		},
	}

	tr := &http.Transport{
		TLSNextProto: map[string]func(authority string, c *tls.Conn) http.RoundTripper{},
	}
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, &http.Client{Transport: tr})
	return Authenticator{
		config:  cfg,
		context: ctx,
	}
}

func (a Authenticator) Token(r *http.Request) (*oauth2.Token, error) {
	values := r.URL.Query()
	if e := values.Get("error"); e != "" {
		return nil, errors.New("spotify: auth failed - " + e)
	}
	code := values.Get("code")
	if code == "" {
		return nil, errors.New("spotify: didn't get access code")
	}

	return a.config.Exchange(a.context, code)
}
