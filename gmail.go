package main

import (
	"log"
	"net/http"

	"google.golang.org/api/gmail/v1"
)

type Email struct {
	Subject string
}

// NewGmailService creates a new service for interacting with Gmail
func NewGmailService(client *http.Client) (*gmail.Service, error) {
	service, err := gmail.New(client)
	if err != nil {
		return nil, err
	}
	return service, nil
}

// FetchRecentEmails fetches the most recent emails up to the specified limit
func FetchRecentEmails(service *gmail.Service, limit int64) ([]Email, error) {
	var emails []Email
	user := "me"
	r, err := service.Users.Messages.List(user).MaxResults(limit).Do()
	if err != nil {
		return nil, err
	}

	for _, m := range r.Messages {
		msg, err := service.Users.Messages.Get(user, m.Id).Format("metadata").MetadataHeaders("Subject").Do()
		if err != nil {
			log.Printf("Unable to retrieve message %v: %v", m.Id, err)
			continue
		}
		for _, header := range msg.Payload.Headers {
			if header.Name == "Subject" {
				emails = append(emails, Email{Subject: header.Value})
				break
			}
		}
	}
	return emails, nil
}
