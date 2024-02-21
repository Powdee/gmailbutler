package main

import (
	"fmt"
	"log"
	"strings"

	"google.golang.org/api/gmail/v1"
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
