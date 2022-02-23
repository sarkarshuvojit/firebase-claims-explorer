/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "firebase-claims-exporer",
	Aliases: []string{"fce"},
	Short:   "A tui application to manage claims in your firebase app",
	Long:    `A tui application to manage claims in your firebase app`,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	var configFilePtr string
	RootCmd.PersistentFlags().StringVarP(
		&configFilePtr,
		"config",
		"c",
		"",
		"Pass config file from firebase admin",
	)
	RootCmd.MarkPersistentFlagRequired("config")
}
