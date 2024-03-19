package main

import (
	"fmt"
	"log"
	"os"
)

func argtest() {
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

	fmt.Println("Good running", args[0])
}

func main() {
	client := GetClient()
	service, err := NewGmailService(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	emails, err := FetchRecentEmails(service, 10)
	if err != nil {
		log.Fatalf("Unable to fetch emails: %v", err)
	}

	for _, email := range emails {
		fmt.Println("Subject:", email.Subject)
	}
}
