/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"firebase.google.com/go/v4/auth"
	"github.com/go-faker/faker/v4"
	"github.com/spf13/cobra"
	"shuvojit.in/firebase-claims-explorer/authentication"
)

var seedCount int

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Used to seed firebase with random users.",
	Long:  "Used to seed firebase with random users.",
	Run: func(cmd *cobra.Command, args []string) {
		configFile, err := cmd.Flags().GetString("config")
		if err != nil {
			fmt.Println("Please set the config flag using --config or -c ")
			os.Exit(1)
		}

		client := authentication.GetAuthClient(configFile)
		seedUsers(seedCount, client)
	},
}

func init() {
	RootCmd.AddCommand(seedCmd)
	seedCmd.Flags().IntVarP(&seedCount, "size", "s", 10, "Number of users to seed. Default is 10.")
}

func resetUserSet(client *auth.Client) {
	savedUsers, _ := authentication.GetAllUsers(client)

	var toBeDeleted []string

	for _, u := range savedUsers {
		toBeDeleted = append(toBeDeleted, u.UID)
	}
	client.DeleteUsers(context.Background(), toBeDeleted)
}

func seedUsers(n int, client *auth.Client) {
	fmt.Printf("Seeding %d users\n", n)
	var users []*auth.UserToImport

	for i := 0; i < n; i++ {
		email := faker.Email()
		users = append(users, (&auth.UserToImport{}).
			Email(email).
			UID(faker.UUIDHyphenated()).
			DisplayName(fmt.Sprintf("%s %s", faker.FirstName(), faker.LastName())).
			CustomClaims(map[string]interface{}{
				"email":    email,
				"birthday": faker.Date(),
				"tz":       faker.Timezone(),
				"currency": faker.Currency(),
				"phone":    faker.Phonenumber(),
			}),
		)
	}

	resetUserSet(client)
	authentication.InsertUsers(client, users)
}
