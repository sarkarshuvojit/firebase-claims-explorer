package authentication

import (
	"context"
	"errors"
	"fmt"
	"log"

	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/iterator"
)

type User struct {
	UID    string
	Email  string
	Claims map[string]interface{}
}

func GetAllUsers(client *auth.Client) ([]User, error) {
	var users []User

	iter := client.Users(context.Background(), "")
	for {
		_user, err := iter.Next()

		if err != nil {
			log.Fatalf("%T\n", err)
		}

		user := User{
			UID:    _user.UID,
			Email:  _user.Email,
			Claims: _user.CustomClaims,
		}

		users = append(users, user)

		if err == iterator.Done {
			break
		}
	}

	return users, nil
}

func SearchUsers(
	client *auth.Client,
	searchFilters []auth.UserIdentifier) ([]*auth.UserRecord, error) {
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
