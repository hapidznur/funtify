package main

import (
	"fmt"
	"log"
	"net/http"
)


func main() {
	http.HandleFunc("/callback/", checkAuth)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})
	// go http.ListenAndServe(":3030", nil)
	url := auth.config.AuthCodeURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)
	if err := http.ListenAndServe(":3030", nil); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

func checkAuth(w http.ResponseWriter, r *http.Request) {
	token, err := auth.Token(r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		fmt.Printf("%q", err)
		log.Fatal(err)
	}

	// use the token to get an authenticated client
	fmt.Printf("%q", token)
	fmt.Fprintf(w, "Login Completed!")
}
