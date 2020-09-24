package main

import (
	"context"
	"crypto/tls"
	"errors"
	"net/http"

	"golang.org/x/oauth2"
)

const (
	// AuthURL is the URL to Spotify Accounts Service's OAuth2 endpoint.
	AuthURL = "https://accounts.spotify.com/authorize"
	// TokenURL is the URL to the Spotify Accounts Service's OAuth2
	// token endpoint.
	TokenURL = "https://accounts.spotify.com/api/token"
)

const (
	// ScopeUserReadPrivate seeks read access to a user's
	// subsription details (type of user account).
	ScopeUserReadPrivate = "user-read-private"
)

type Authenticator struct {
	config  *oauth2.Config
	context context.Context
}

func Authorize(redirectURL string, scopes ...string) Authenticator {
	// ini comment
	// test comment
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

// Token pulls an authorization code from an HTTP request and attempts to exchange
// it for an access token.  The standard use case is to call Token from the handler
// that handles requests to your application's redirect URL.
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
