package backend

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

var MailSrv *gmail.Service

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
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

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func sendConfirmationEmail(name string, email string) {

	// Create email message in HTML format
	fromLine := "From: hacknjit@njit.edu\r\n"
	toLine := fmt.Sprintf("To: %s\r\n", email)
	subjectLine := "Subject: HackNJIT 2025 Registration Confirmation\r\n"
	headerLine := "Content-Type: text/html; charset=UTF-8\r\n\r\n"
	messageBody := fmt.Sprintf("<html><body><p>Hello %s,</p><p>Thank you for registering for HackNJIT!</p><p>This email serves as confirmation of your registration. There is nothing further you need to do right now; all updates will be on the HackNJIT website, hacknjit.org, and we will continue to send updates via email before the event. We will also email you again closer to the new date with a guide to attending in-person.</p><p>Thank you again, and we look forward to seeing you at HackNJIT on <b>March 7-8th, 2026</b>!</p><p>Sincerely,<br/>The HackNJIT Team</p></body></html>", name)
	messageStr := []byte(fromLine + toLine + subjectLine + headerLine +
		messageBody)

	// Encode message in base64 URL-safe format
	message := gmail.Message{
		Raw: base64.URLEncoding.EncodeToString(messageStr),
	}

	_, err := MailSrv.Users.Messages.Send("me", &message).Do()
	if err != nil {
		log.Fatalf("Unable to send email: %v", err)
	} else {
		log.Printf("[DEBUG] Email successfully send to %s\n", email)
	}
}

func init() {
	log.Println("Enabling email service...")
	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	log.Println("\tClient secrets read...")

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailSendScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)
	log.Println("\tHTTP Client configured...")

	MailSrv, err = gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}
	log.Println("\tMail Client Successfully Configured")
}
