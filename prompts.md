# Prompts
This file lists the GPT prmopts I used to learn and implement this project. The ones with <-- identify the most useful prompts/results.

## Getting started
1. i want to build a local-first go program that can be run from the command line.
i want to connect to my gmail account and create a local database of all the emails i have received. if the gmail connection requires authentication, i want to be able to enter my credentials via interactive prompt in the command line.

2. you are a golang developer, and you want to build a local-first go program that can be run from the command line from scratch. you start with an empty folder. you want a user to be able to connect to their gmail account and fetch the contents and metadata of the last 10 emails they have received. if the gmail connection requires authentication, you want to be able to enter their credentials via interactive prompt in the command line. organize the functionality into separate go files, and list the dependencies in a go.mod file. show me the code for each file with the appropriate filename, and explain any steps needed to be taken to complete the set up of the project to print the subject line of each email fetched in the console. <---

## Building on top
1. modify auth.go in the following ways:
- add a check to GetClient to check if the token is credentials are valid and the token is not expired
- when getTokenFromWeb is called, initiate a web server in a separate thread to listen for the callback from the browser with the auth code
- parse the auth code from that callback, print it to the console, and then use it to complete the token refresh

2. modify gmail.go in the following ways:
- add a field "tags" to the Email struct
- modify FetchRecentEmails to add one or more of the following tags if applicable to to each email in the returned slice:
    - "#unread"
    - "#archived"
    - "#promotion"
    - "#update"
    - "#social"
