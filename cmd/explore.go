package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"firebase.google.com/go/v4/auth"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"shuvojit.in/firebase-claims-explorer/authentication"
	"shuvojit.in/firebase-claims-explorer/components"
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
	users, err := authentication.GetAllUsers(client)

	if err != nil {
		panic(err)
	}

	clearScreen()

	p := tea.NewProgram(createExploreModel(users))
	if err := p.Start(); err != nil {
		clearScreen()
		fmt.Printf("Some error occored. %v", err)
		os.Exit(1)
	}
}

func clearScreen() {
	clearcmd := exec.Command("clear")
	clearcmd.Stdout = os.Stdout
	clearcmd.Run()
}

type exploreModel struct {
	users             []authentication.User
	selectedUserIndex int
	selectedUser      authentication.User

	searchQuery     string
	filteredResults []authentication.User
}

func createExploreModel(users []authentication.User) exploreModel {

	return exploreModel{
		users:             users,
		selectedUserIndex: 0,

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
			clearScreen()
			return m, tea.Quit
		case "j", "down":
			m.MoveCursorDown()
			break
		case "k", "up":
			m.MoveCursorUp()
			break
		}
	}

	return m, nil
}

func (m *exploreModel) MoveCursorDown() {
	if m.selectedUserIndex == len(m.users)-1 {
		return
	}
	m.selectedUserIndex += 1
}

func (m *exploreModel) MoveCursorUp() {
	if m.selectedUserIndex == 0 {
		return
	}
	m.selectedUserIndex -= 1
}

func (m exploreModel) View() string {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA"))

	output := components.UserList(m.users, m.users[m.selectedUserIndex])
	return style.Render(output)
}
