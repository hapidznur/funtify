package main

import "testing"

func TestAuthUser(t *testing.T) {
	got := AuthUser()
	want := "a"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
