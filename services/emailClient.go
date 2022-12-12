package services

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"encoding/base64"

	"github.com/S3B4SZ17/Web_Algo/management"
	"go.uber.org/zap"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

// func ConfigureEmailService() *http.Client {

// 	b, err := os.ReadFile("credentials.json")
// 	if err != nil {
// 		management.Log.Error("Unable to read client secret file:", zap.String("credentials_file", err.Error()))
// 		os.Exit(1)
// 	}

// 	// If modifying these scopes, delete your previously saved token.json.
// 	config, err := google.ConfigFromJSON(b, gmail.GmailReadonlyScope)
// 	if err != nil {
// 		management.Log.Error("Unable to parse client secret file to config:", zap.String("err_message", err.Error()))
// 		os.Exit(1)
// 	}

// 	client := getClient(config)

// 	return client

// }

func GetGmailService(ctx context.Context, client *http.Client) (*gmail.Service, error) {

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		management.Log.Error("Unable to retrieve Gmail client:", zap.String("err_message", err.Error()))
		os.Exit(1)
	}

	return srv, err
}

func GetGmailLabels(mail *gmail.Service) *EmailResponse {

	// Hardcoding me since it represents the context of the user that is authenticated.
	// Reference this thread https://stackoverflow.com/questions/26135310/gmail-api-returns-403-error-code-and-delegation-denied-for-user-email
	user := "me"
	r, err := mail.Users.Labels.List(user).Do()
	if err != nil {
		management.Log.Error("Unable to retrieve labels:", zap.String("err_message", err.Error()))
		os.Exit(1)
	}
	if len(r.Labels) == 0 {
		fmt.Println("No labels found.")
		return nil
	}
	fmt.Println("Labels:")
	var labels string
	for _, l := range r.Labels {
		labels += fmt.Sprintf("- %s\n", l.Name)
	}

	response := &EmailResponse{Labels: labels}
	return response
}

func SendGmailEmail(mail *gmail.Service, email_to *EmailTo) (*EmailResponse, error) {

	// Hardcoding me since it represents the context of the user that is authenticated.
	// Reference this thread https://stackoverflow.com/questions/26135310/gmail-api-returns-403-error-code-and-delegation-denied-for-user-email
	user := "me"

	// to := &gmail.MessagePartHeader{Name: "To", Value: email_to.To}
	// listHeaders := []*gmail.MessagePartHeader{to}
	// sEnc := b64.URLEncoding.EncodeToString([]byte(email_to.Body))
	mail_text := fmt.Sprintf("From: sebaszh17@gmail.com\nTo: %s\nSubject: New Algo resolved!!\n\n %s", email_to.To, email_to.Body)

	raw := base64.URLEncoding.EncodeToString([]byte(mail_text))
	// ms_body := &gmail.MessagePartBody{Data: sEnc}
	// ms_part := &gmail.MessagePart{Headers: listHeaders, Body: ms_body}
	emailMessage := &gmail.Message{Raw: raw}
	_, err := mail.Users.Messages.Send(user, emailMessage).Do()

	if err != nil {
		err := errors.New(err.Error())
		return nil, err
	}

	response := &EmailResponse{Message: "Email sent successfully"}
	return response, nil
}

type EmailTo struct {
	From string `json:"from"`
	To   string `json:"to"`
	Body string `json:"body"`
}
type EmailResponse struct {
	Labels  string `json:"labels"`
	Message string `json:"message,omitempty"`
}
