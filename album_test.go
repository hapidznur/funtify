package main

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
