/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"firebase.google.com/go/v4/auth"
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

func seedUsers(n int, client *auth.Client) {
	fmt.Printf("Seeding %d users\n", n)
	var users []*auth.UserToImport

	for i := 0; i < n; i++ {
		users = append(users, (&auth.UserToImport{}).
			UID(fmt.Sprintf("UID%d", i)).
			DisplayName(fmt.Sprintf("Name%d", i)).
			CustomClaims(map[string]interface{}{"index": i}),
		)
	}

	authentication.InsertUsers(client, users)
}
