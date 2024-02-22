package main

import (
	"google.golang.org/api/gmail/v1"
)

type GmailService interface {
	ListMessages(userId string, query string, maxResults int64, pageToken string) (*gmail.ListMessagesResponse, error)
	BatchDeleteMessages(userId string, ids []string) error
}

type gmailServiceWrapper struct {
	srv *gmail.Service
}

func NewGmailServiceWrapper(srv *gmail.Service) GmailService {
	return &gmailServiceWrapper{srv: srv}
}

func (g *gmailServiceWrapper) ListMessages(userId string, query string, maxResults int64, pageToken string) (*gmail.ListMessagesResponse, error) {
	return g.srv.Users.Messages.List(userId).Q(query).MaxResults(maxResults).PageToken(pageToken).Do()
}

func (g *gmailServiceWrapper) BatchDeleteMessages(userId string, ids []string) error {
	return g.srv.Users.Messages.BatchDelete(userId, &gmail.BatchDeleteMessagesRequest{Ids: ids}).Do()
}
