package main

import (
	"log"
	"net/http"

	"google.golang.org/api/gmail/v1"
)

type Email struct {
	Subject string
	Tags    []string
}

// contains checks if a slice contains a specific string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
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
		email := Email{}
		for _, header := range msg.Payload.Headers {
			if header.Name == "Subject" {
				email.Subject = header.Value
				break
			}
		}

		for _, label := range msg.LabelIds {
			switch label {
			case "UNREAD":
				email.Tags = append(email.Tags, "#unread")
			case "CATEGORY_PROMOTIONS":
				email.Tags = append(email.Tags, "#promotion")
			case "CATEGORY_UPDATES":
				email.Tags = append(email.Tags, "#update")
			case "CATEGORY_SOCIAL":
				email.Tags = append(email.Tags, "#social")
			}
		}
		if !contains(msg.LabelIds, "INBOX") {
			email.Tags = append(email.Tags, "#archived")
		}
		emails = append(emails, email)
	}

	return emails, nil
}
