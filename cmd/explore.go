package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"firebase.google.com/go/v4/auth"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"shuvojit.in/firebase-claims-explorer/authentication"
)

var globalClient *auth.Client
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
		globalClient = client
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

	p := tea.NewProgram(createExploreModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Some error occored. %v", err)
		os.Exit(1)
	}
}

type exploreModel struct {
	users        []authentication.User
	selectedUser authentication.User

	searchQuery     string
	filteredResults []authentication.User
}

func createExploreModel() exploreModel {
	users, err := authentication.GetAllUsers(globalClient)
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

func (m exploreModel) Init() tea.Cmd {
	return nil
}

func (m exploreModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
		// Update model
	}

	return m, nil
}

func (m exploreModel) View() string {
	return "Something is printed, innit?"
}
