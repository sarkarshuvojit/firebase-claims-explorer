/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "firebase-claims-exporer",
	Short: "A tui application to manage claims in your firebase app",
	Long:  `A tui application to manage claims in your firebase app`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	var configFilePtr string
	rootCmd.PersistentFlags().StringVarP(
		&configFilePtr,
		"config",
		"c",
		"",
		"Pass config file from firebase admin",
	)
	rootCmd.MarkPersistentFlagRequired("config")
}
