package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	args := os.Args[1:] // This will include the `--`, so we might want to remove it.

	// Detect if there is a `--` and remove it and everything before it
	for i, arg := range args {
		if arg == "--" {
			args = args[i+1:]
			break
		}
	}

	if len(args) < 1 {
		fmt.Println("Usage: gmpress <path>")
		fmt.Println("[or] go run gmpress -- <path>")
		os.Exit(1)
	}

	fmt.Println("Using cred file:", args[0])

	client := GetClient(args[0])
	service, err := NewGmailService(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	} else {
		fmt.Println("...connected to Gmail")
	}

	emails, err := FetchRecentEmails(service, 10)
	if err != nil {
		log.Fatalf("Unable to fetch emails: %v", err)
	}

	for _, email := range emails {
		fmt.Println("Subject:", email.Subject, " tags: ", email.Tags)
	}
}
