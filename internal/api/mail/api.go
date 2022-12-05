package mail

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"google.golang.org/api/gmail/v1"
	"log"
	"net/http"
	"os"
)

func TokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)

	return tok, err
}

func SaveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()

	json.NewEncoder(f).Encode(token)
}

func GetClient(config *oauth2.Config) *http.Client {
	tokFile := "token.json"

	tok, err := TokenFromFile(tokFile)
	if err != nil {
		tok = GetTokenFromWeb(config)
		SaveToken(tokFile, tok)
	}

	return config.Client(context.Background(), tok)
}

func GetTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authUrl := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	fmt.Println("Go to the following link and type the authorization code:\n", authUrl)

	var authCode string

	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalln("Unable to read auth code:", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalln("Unable to retrieve token:", err)
	}

	return tok
}

func SendMail(from string, to, title, message interface{}, srv *gmail.Service) (bool, error) {
	msgStr := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", from, to, title, message)
	msg := []byte(msgStr)

	gMessage := &gmail.Message{Raw: base64.URLEncoding.EncodeToString(msg)}

	_, err := srv.Users.Messages.Send("me", gMessage).Do()
	if err != nil {
		fmt.Println("Could not send mail>", err)
		return false, err
	}

	return true, nil
}
