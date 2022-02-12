package auth

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func GetApp(configFile string) *firebase.App {
	opt := option.WithCredentialsFile(configFile)
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		panic(err)
	}

	return app
}
