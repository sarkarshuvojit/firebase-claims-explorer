package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

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

func main() {
	abspath := getConfigFilePath()
	client := auth.GetAuthClient(abspath)
	fmt.Printf("Client: %v\n", client)
}
