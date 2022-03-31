// The authentication uses the authorization code flow with PKCE(RFC 7636). It provides protection
// against attacks where the authorization code may be intercepted.

package main

import (
	"context"
	"fmt"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const redirectURI = "http://localhost:8080/callback"

type LoginSession struct {
	auth *spotifyauth.Authenticator
	URL  string
}

type PixyConfig struct {
	CodeVerifier
}

type EnvConfig struct {
	ClientID     string
	ClientSecret string
	BaseURL      string
}

// TODO: Clean up the global vars and their init.
var (
	ch    = make(chan *spotify.Client)
	state = getState(10)

	_ = os.Setenv("SPOTIFY_ID", "75b4e7029ff846419fef8ad0e7d86231")
	_ = os.Setenv("SPOTIFY_SECRET", "eed9e286ec6a4b62b09b2e3515ab08d5")

	clientID      = os.Getenv("SPOTIFY_ID")
	clientSecret  = os.Getenv("SPOTIFY_SECRET")
	cv, _         = CreateCodeVerifier()
	codeVerifier  = cv.Value
	codeChallenge = cv.CodeChallengeS256()
)

func loginSession() {
	authServer()
	auth := newAuth()
	URL := authURL(auth)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", URL)

	client := <-ch
	user, err := client.CurrentUser(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("You are logged in as:", user.ID)
}

func authServer() {
	http.HandleFunc("/callback", completeAuth)
	go http.ListenAndServe(":8080", nil)
}

func newAuth() *spotifyauth.Authenticator {
	auth := spotifyauth.New(
		spotifyauth.WithClientID(clientID),
		spotifyauth.WithClientSecret(clientSecret),
		spotifyauth.WithRedirectURL(redirectURI),
		spotifyauth.WithScopes(
			spotifyauth.ScopeUserReadPrivate,
			spotifyauth.ScopeUserReadCurrentlyPlaying,
			spotifyauth.ScopeUserReadPlaybackState,
			spotifyauth.ScopeUserReadRecentlyPlayed,
		))
	return auth
}

func authURL(auth *spotifyauth.Authenticator) string {
	url := auth.AuthURL(
		state,
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
		oauth2.SetAuthURLParam("code_challenge", codeChallenge),
	)

	return url
}

func getState(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}

func main() {

	//// first start an HTTP server
	//http.HandleFunc("/callback", completeAuth)
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	log.Println("Got request for:", r.URL.String())
	//})
	//go http.ListenAndServe(":8080", nil)
	//
	//url := auth.AuthURL(state,
	//	oauth2.SetAuthURLParam("code_challenge_method", "S256"),
	//	oauth2.SetAuthURLParam("code_challenge", codeChallenge),
	//)
	//fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)
	//
	//// wait for auth to complete
	//client := <-ch
	//
	//// use the client to make calls that require authorization
	//user, err := client.CurrentUser(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("You are logged in as:", user.ID)

}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	//tok, err := auth.Token(r.Context(), state, r,
	//	oauth2.SetAuthURLParam("code_verifier", codeVerifier))
	//if err != nil {
	//	http.Error(w, "Couldn't get token", http.StatusForbidden)
	//	log.Fatal(err)
	//}
	//if st := r.FormValue("state"); st != state {
	//	http.NotFound(w, r)
	//	log.Fatalf("State mismatch: %s != %s\n", st, state)
	//}
	//// use the token to get an authenticated client
	//client := spotify.New(auth.Client(r.Context(), tok))
	//fmt.Fprintf(w, "Login Completed!")
	//ch <- client
}
