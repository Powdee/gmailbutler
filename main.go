package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func deleteUnreadEmails(srv *gmail.Service, userId string) {
	var allUnreadMessageIds []string
	nextPageToken := ""
	loadingMessage := "Loading unread emails"
	dots := ""

	for {
		fmt.Printf("\r%s%s   ", loadingMessage, dots) // Print loading message with dynamic dots
		r, err := srv.Users.Messages.List(userId).Q("is:unread").MaxResults(500).PageToken(nextPageToken).Do()
		if err != nil {
			log.Fatalf("Unable to retrieve unread emails: %v", err)
		}

		for _, m := range r.Messages {
			allUnreadMessageIds = append(allUnreadMessageIds, m.Id)
		}

		nextPageToken = r.NextPageToken
		if nextPageToken == "" {
			break 
		}

		if len(dots) < 3 {
			dots += "."
		} else {
			dots = ""
		}
	}

	fmt.Printf("\r%s\r", strings.Repeat(" ", len(loadingMessage+dots+"   "))) // Clear the loading message line

	if len(allUnreadMessageIds) == 0 {
		fmt.Println("No unread emails to delete.")
		return
	}

	err := srv.Users.Messages.BatchDelete(userId, &gmail.BatchDeleteMessagesRequest{Ids: allUnreadMessageIds}).Do()
	if err != nil {
		log.Fatalf("Unable to delete unread emails: %v", err)
	}

	fmt.Printf("Deleted %d unread emails.\n", len(allUnreadMessageIds))
}

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

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func main() {

	ctx := context.Background()
	b, err := os.ReadFile(".secrets/credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	user := "me"
	deleteUnreadEmails(srv, user)
}