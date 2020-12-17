package funtify

import (
	"net/http"
	"testing"
)

// FInd ALbum
func TestFindAlbum(t *testing.T) {
	server := testServerFromFile(http.StatusOK, "tests_data/find_album.txt")
	client := testClient(server)

	defer server.Close()

	album, err := client.GetAlbum(ID("0sNOF9WDwhWunNAHPD3Baj"))

	if err != nil {
		t.Fatal(err)
	}

	if album == nil {
		t.Fatal("Got nil Album")
	}

	if album.Name != "She's So Unusual" {
		t.Error("Got Wrong Album")
	}

	released := album.ReleaseDateTime()

	if released.Year() != 1983 {
		t.Errorf("Expeted 1983, got %d\n", released.Year())
	}
}

// Test Get Album Tracks
func TestFindAlbumTracks(t *testing.T) {
	server := testServerFromFile(http.StatusOK, "tests_data/find_album_tracks.txt")
	client := testClient(server)

	defer server.Close()

	result, err := client.GetAlbumTracks(ID("0sNOF9WDwhWunNAHPD3Baj"), 1, 0)

	if err != nil {
		t.Fatal(err)
	}

	if result.Count != 13 {
		t.Fatal("Got", result.Count, "want 13")
	}

	if len(result.Tracks) == 1 {
		if result.Tracks[0].Name != "Money Changes Everything" {
			t.Error("Expected track 'Money Changes Everything', got", result.Tracks[0].Name)
		}
	} else {
		t.Error("Expected 1 track, got", len(result.Tracks))
	}
}
