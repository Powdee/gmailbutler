package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

func getClient(config *oauth2.Config) *http.Client {
	tokFile := ".secrets/token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
		fmt.Println("Authorize this app by visiting this url:", authURL)
		http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
			code := r.URL.Query().Get("code")
			if code == "" {
				fmt.Fprintf(w, "Authorization code not found in the request")
				return
			}
			tok, err := config.Exchange(context.Background(), code)
			if err != nil {
				log.Fatalf("Unable to retrieve token from web: %v", err)
			}
			saveToken(tokFile, tok)
			fmt.Fprintf(w, "Authorization successful, you can now close this window.")
		})
		log.Fatal(http.ListenAndServe(":8080", nil))
		return nil
	}
	return config.Client(context.Background(), tok)
}
