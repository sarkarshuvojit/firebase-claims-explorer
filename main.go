package main

import (
	"github.com/spf13/cobra/doc"
	"shuvojit.in/firebase-claims-explorer/cmd"
)

func main() {
	err := doc.GenMarkdownTree(cmd.RootCmd, "./docs")
	if err != nil {
		panic(err)
	}
	cmd.Execute()
}
