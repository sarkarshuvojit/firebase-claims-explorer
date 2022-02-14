package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	fb "firebase.google.com/go/v4/auth"
	"google.golang.org/api/iterator"
	"shuvojit.in/firebase-claims-exporer/auth"
)

func getConfigFilePath() string {
	configFilePtr := flag.String("config", "default", "path to your firebase config")

	flag.Parse()

	if *configFilePtr == "default" {
		fmt.Println("Please specify a config file using config flag")
		os.Exit(1)
	}

	if !strings.HasSuffix(*configFilePtr, ".json") {
		fmt.Println("Please specify valid config file")
		os.Exit(1)
	}

	return *configFilePtr
}

func listScreen(client *fb.Client) error {
	iter := client.Users(context.Background(), "")
	idx := 0
	for {
		user, err := iter.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			return err
		}

		fmt.Printf("%d, %s", idx, user.Email)
		idx++
	}
	return nil
}

func start() error {
	abspath := getConfigFilePath()
	client := auth.GetAuthClient(abspath)
	err := listScreen(client)

	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := start()
	if err != nil {
		panic("Oops, something broke")
	}
}
