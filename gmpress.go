package main

import (
	"fmt"
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

	fmt.Println("Good running", args[0])
}
