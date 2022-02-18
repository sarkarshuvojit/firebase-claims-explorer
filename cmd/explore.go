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

// exploreCmd represents the explore command
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
}
