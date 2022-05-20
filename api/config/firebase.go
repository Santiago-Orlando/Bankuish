package config

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"

	"google.golang.org/api/option"
)

func Firebase() (*auth.Client, error) {

	opt := option.WithCredentialsFile("api/config/bankuish-a44ad-firebase-adminsdk-81ezn-9c3a9b8171.json")
	config := &firebase.Config{ProjectID: "bankuish-a44ad"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		return nil, err
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	return client, nil
}
