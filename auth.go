package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

const (
	tokenFile = "token.json"
)

var (
	mu               sync.Mutex
	authCodeResponse = make(chan string)
)

// GetClient retrieves a token, saves the token, then returns the generated client.
func GetClient(credentialPath string) *http.Client {
	// Load client secrets
	if credentialPath == "" {
		credentialPath = "../gmpress-aux/credentials.json"
	}
	cred_json, err := os.ReadFile(credentialPath)
	// Try gain with default path
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(cred_json, gmail.GmailReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)
	return client
}

func getClient(config *oauth2.Config) *http.Client {
	token, err := tokenFromFile(tokenFile)

	if err != nil {
		log.Printf("Unable to read token file: %v", err)
	} else if !token.Valid() {
		log.Println("Token is invalid, refreshing")
	}

	if err != nil || !token.Valid() {
		log.Println("Going to get token from web, human interaction required...")
		token, err = getTokenFromWeb(config)
		if err != nil {
			log.Fatalf("Unable to refresh token: %v", err)
		} else {
			saveToken(tokenFile, token)
		}
	}
	return config.Client(context.Background(), token)
}

// getTokenFromWeb requests a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {

	go startAuthCodeServer(config)

	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	log.Printf("Go to the following link in your browser and continue through the prompts: \n%v\n", authURL)

	authCode := <-authCodeResponse
	fmt.Printf("Auth code received from response: %s\n", authCode)
	token, err := config.Exchange(context.Background(), authCode)
	return token, err
}

// startServer starts a web server that listens on http://localhost:8080/callback
func startAuthCodeServer(config *oauth2.Config) {
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()

		code := r.URL.Query().Get("code")
		fmt.Fprintf(w, "Authorization received. You can now close this browser tab.")
		fmt.Printf("Param: %s\n", code)
		authCodeResponse <- code
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// tokenFromFile retrieves a token from a given file path.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	token := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(token)
	return token, err
}

// saveToken saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	log.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache OAuth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
