package auth

import (
	"context"
	"errors"

	"firebase.google.com/go/v4/auth"
)

func SearchUsers(client *auth.Client, searchFilters []auth.UserIdentifier) ([]*auth.UserRecord, error) {
	result, err := client.GetUsers(context.Background(), searchFilters)

	if err != nil {
		return nil, errors.New("Failed to get users")
	}

	return result.Users, nil
}
