/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"firebase.google.com/go/v4/auth"
	"github.com/spf13/cobra"
	"shuvojit.in/firebase-claims-explorer/authentication"
)

var exploreCmd = &cobra.Command{
	Use:   "explore",
	Short: "Runs tui app to list users and view claims",
	Long:  "Runs tui app to list users and view claims",
	Run: func(cmd *cobra.Command, args []string) {
		configFile, err := cmd.Flags().GetString("config")
		if err != nil {
			fmt.Println("Please set the config flag using --config or -c ")
			os.Exit(1)
		}

		client := authentication.GetAuthClient(configFile)
		launchTui(client)
	},
}

func init() {
	rootCmd.AddCommand(exploreCmd)
}

func launchTui(client *auth.Client) {
	fmt.Printf("TUI needs to be invoked here\n")
	users, err := authentication.GetAllUsers(client)

	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(users)
	fmt.Println(string(b))
}

type exploreModel struct {
	users        []authentication.User
	selectedUser authentication.User

	searchQuery     string
	filteredResults []authentication.User
}

func createExploreModel(client *auth.Client) exploreModel {
	users, err := authentication.GetAllUsers(client)
	if err != nil {
		panic(err)
	}

	return exploreModel{
		users:        users,
		selectedUser: authentication.User{},

		searchQuery:     "",
		filteredResults: []authentication.User{},
	}
}
