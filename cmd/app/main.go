package main

import (
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"io/ioutil"
	"log"
	"webapp/internal/api/mail"
	"webapp/internal/database"
	"webapp/internal/user"
)

var from = "REPLACE IT!"

func main() {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, gmail.GmailSendScope)
	if err != nil {
		log.Fatalf("unable to parse client secret file config: %v", err)
	}

	client := mail.GetClient(config)

	srv, err := gmail.New(client)
	if err != nil {
		log.Fatalf("unable to retrieve Gmail client: %v", err)
	}

	to, title, message, err := user.AskUser()
	if err != nil {
		log.Fatalf("unable to retrieve delivery email: %v", err)
	}

	_, err = mail.SendMail(from, to, title, message, srv)
	if err != nil {
		log.Fatalf("error while sending your message. Please check delivery email: %v", err)
	}

	nowTime, err := user.Output(message, title, to)
	if err != nil {
		log.Fatalf("error while displaying message to user: %v", err)
	}

	dbClient, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	if err != nil {
		log.Fatalf("error while connecting to database: %v", err)
	}
	defer database.Close(dbClient, ctx, cancel)

	var doc interface{}

	doc = bson.D{
		{"From", from},
		{"To", to},
		{"Title", title},
		{"Message", message},
		{"Time", nowTime},
	}

	_, err = database.InsertOne(dbClient, ctx, "Mails", "data", doc)
	if err != nil {
		log.Fatalf("error while inserting data to mongo: %v", err)
	}
}
