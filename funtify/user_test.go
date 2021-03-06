package funtify

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func testServer(code int, body io.Reader) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		io.Copy(w, body)
	}))

	return server
}

func testServerFromFile(code int, filename string) *httptest.Server {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return testServer(code, f)
}

func testClient(server *httptest.Server) *Client {
	client := &Client{
		http:    http.DefaultClient,
		baseURL: server.URL + "/",
	}
	return client
}

const userResponse = `
{
  "display_name" : "Ronald Pompa",
  "external_urls" : {
    "spotify" : "https://open.spotify.com/user/wizzler"
    },
    "followers" : {
      "href" : null,
      "total" : 3829
    },
    "href" : "https://api.spotify.com/v1/users/wizzler",
    "id" : "wizzler",
    "images" : [ {
      "height" : null,
      "url" : "http://profile-images.scdn.co/images/userprofile/default/9d51820e73667ea5f1e97ea601cf0593b558050e",
      "width" : null
    } ],
    "type" : "user",
    "uri" : "spotify:user:wizzler"
}`

const currentUser = `{
	"country" : "US",
	"display_name" : null,
	"email" : "username@domain.com",
	"external_urls" : {
		"spotify" : "https://open.spotify.com/user/username"
	},
	"followers" : {
		"href" : null,
		"total" : 0
	},
	"href" : "https://api.spotify.com/v1/users/userame",
	"id" : "username",
	"images" : [ ],
	"product" : "premium",
	"type" : "user",
	"uri" : "spotify:user:username",
	"birthdate" : "1985-05-01"
}`

func TestUserCurrent(t *testing.T) {
	server := testServer(http.StatusOK, strings.NewReader(currentUser))
	defer server.Close()

	client := testClient(server)
	me, errr := client.CurrentUser()
	if errr != nil {
		t.Error(errr)
		return
	}
	if me.Product != "premium" {
		t.Errorf("got %s, want %s", me.Product, "premium")
	}
}

func TestUserProfile(t *testing.T) {
	server := testServer(http.StatusOK, strings.NewReader(userResponse))

	client := testClient(server)
	userID := "Ronald Pompa"
	profile, err := client.UserProfile(userID)

	assertError(t, err)
	if profile.DisplayName != userID {
		t.Errorf("got %s, want %s", profile.DisplayName, userID)
	}
}

func TestCurrentUserTrack(t *testing.T) {
	filename := "tests_data/current_users_tracks.txt"
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	server := testServer(http.StatusOK, f)
	client := testClient(server)
	track, err := client.CurrentUserTracks()
	assertError(t, err)
	assertTrue(t, track.Endpoint, "https://api.spotify.com/v1/me/tracks?offset=0&limit=20")
}

func assertError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Error %s", err)
	}
}

func assertTrue(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
