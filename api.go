package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	redirectURL = readEnvVariable("redirect_url")
	auth        = Authorize(redirectURL, ScopeUserReadPrivate)
	ch    = make(chan *Client)
	state       = "superstring"
)

// AuthUser for API get AuthUser from spotify
func AuthUser() {
	http.HandleFunc("/callback", completeAuth)
	// wait for auth to complete
	client := <-ch
	user,err := client.CurrentUser()
	fmt.Print(err)
	fmt.Print(user)
}

func Authenticate(token string) string {
	return "success"
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}
	// use the token to get an authenticated client
	client := auth.NewClient(tok)
	fmt.Fprintf(w, "Login Completed!")
	ch <- &client
}
