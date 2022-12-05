# Project objective
> The aim of the project is to gain hands-on experience of using the GMAIL API and synchronously entering the resulting data into the database.
# Installation
> Write in terminal: `git clone https://github.com/5t4lk/mail-api`.
# How to run
> Enter to folder named "webapp". Write in terminal: `go run .\cmd\app\main.go`.
# Description
> The program allows the user to save, modify or delete the information from database that they enter.
# Correct use of the service
- Step 1. In the [Google Cloud console](https://console.cloud.google.com/), go to Menu menu > APIs & Services > Credentials.
- - Click Create Credentials > OAuth client ID.
- - Click Application type > Web application.
- - In the "Name" field, type a name for the credential. This name is only shown in the Google Cloud console.
- - Click Create. The OAuth client created screen appears, showing your new Client ID and Client secret. Download it and save as "credentials.json". Move "credentials.json" to your cloned directory "webapp"
- - Click OK. The newly created credential appears under "OAuth 2.0 Client IDs."
- Step 2. Write these 2 strings to your console:
- - >go get google.golang.org/api/gmail/v1
- - >go get golang.org/x/oauth2/google
- Step 3. Write in console:
- - >go run .\cmd\app\main.go
- - The first time you run the sample, it prompts you to authorize access:
- - - If you're not already signed in to your Google Account, you're prompted to sign in. If you're signed in to multiple accounts, select one account to use for authorization.
- - - Click Accept.
- - - Copy the code from the browser, paste it into the command-line prompt, and press Enter.
- Authorization information is stored in the file system as "token.json", so the next time you run the sample code, you aren't prompted for authorization.
- You have successfully used my service that makes requests to the Gmail API.




# Stack
- Golang
- API
- MongoDB
- Git