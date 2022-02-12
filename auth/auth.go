package auth

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	auth "firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

func getApp(configFile string) *firebase.App {
	opt := option.WithCredentialsFile(configFile)
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		log.Panic("Error initialising firebase app")
		panic(err)
	}

	return app
}

func GetAuthClient(configFile string) *auth.Client {
	app := getApp(configFile)
	client, err := app.Auth(context.Background())

	if err != nil {
		log.Panic("Error initialising auth")
		panic(err)
	}

	return client
}
