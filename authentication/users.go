package authentication

import (
	"context"
	"errors"
	"fmt"

	"firebase.google.com/go/v4/auth"
)

func SearchUsers(client *auth.Client, searchFilters []auth.UserIdentifier) ([]*auth.UserRecord, error) {
	result, err := client.GetUsers(context.Background(), searchFilters)

	if err != nil {
		return nil, errors.New("Failed to get users")
	}

	return result.Users, nil
}

func InsertUsers(client *auth.Client, users []*auth.UserToImport) error {
	result, err := client.ImportUsers(context.Background(), users)
	if err != nil {
		return err
	}
	fmt.Printf("Seeding result: %v\n", result)
	return nil
}
