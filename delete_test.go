package main

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"google.golang.org/api/gmail/v1"
)

type MockGmailService struct {
	mock.Mock
}

func (m *MockGmailService) ListMessages(userId string, query string, maxResults int64, pageToken string) (*gmail.ListMessagesResponse, error) {
	args := m.Called(userId, query, maxResults, pageToken)
	return args.Get(0).(*gmail.ListMessagesResponse), args.Error(1)
}

func (m *MockGmailService) BatchDeleteMessages(userId string, ids []string) error {
	args := m.Called(userId, ids)
	return args.Error(0)
}

func TestDeleteUnreadEmails(t *testing.T) {
	mockGmailService := new(MockGmailService)
	mockResp := &gmail.ListMessagesResponse{
		Messages: []*gmail.Message{
			{Id: "id1"},
			{Id: "id2"},
		},
	}

	mockGmailService.On("ListMessages", "me", "is:unread", int64(500), "").Return(mockResp, nil)
	mockGmailService.On("BatchDeleteMessages", "me", []string{"id1", "id2"}).Return(nil)

	deleteUnreadEmails(mockGmailService, "me")

	mockGmailService.AssertExpectations(t)
}
