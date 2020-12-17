package funtify

import (
	"net/http"
	"testing"
)

func TestRecentlyPlayed(t *testing.T) {
	server := testServerFromFile(http.StatusOK, "tests_data/player_recently_played.txt")
	client := testClient(server)
	recently, err := client.GetRecentlyPlayed()
	assertError(t, err)

	if len(recently.Items) != 20 {
		t.Error("Too few or too many items were returned")
	}

	actualTimePhrase := recently.Items[0].PlayedAt.Format("2006-01-02T15:04:05.999Z")
	expectedTimePhrase := "2017-05-27T20:07:54.721Z"
	assertTrue(t, actualTimePhrase, expectedTimePhrase)
}
