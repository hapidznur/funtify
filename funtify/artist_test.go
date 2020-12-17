package funtify

import (
	"net/http"
	"testing"
)

func TestFindArtist(t *testing.T) {
	server := testServerFromFile(http.StatusOK, "tests_data/find_artist.txt")
	client := testClient(server)

	artist, err := client.GetArtist(ID("0TnOYISbd1XYRBk9myaseg"))

	if err != nil {
		t.Fatal(err)
	}

	if followers := artist.Followers.Total; followers != 2265279 {
		t.Errorf("got %d followers, want 226579\n", followers)
	}

	if artist.Name != "Pitbull" {
		t.Error("GOt ", artist.Name, ", wanted Pitbull")
	}
}
