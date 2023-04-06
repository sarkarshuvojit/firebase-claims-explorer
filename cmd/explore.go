package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"firebase.google.com/go/v4/auth"
	"github.com/charmbracelet/bubbles/textarea"
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
	RootCmd.AddCommand(exploreCmd)
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

type ScreenRender func(m exploreModel) string

type Screen struct {
	Name   string
	Render ScreenRender
}

var (
	LIST_SCREEN = Screen{
		Name: "LIST_SCREEN",
		Render: func(m exploreModel) string {
			var style = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#FAFAFA"))

			output := components.UserList(m.users, m.users[m.selectedUserIndex])
			return style.Render(output)
		}}
	DETAIL_SCREEN = Screen{
		Name: "DETAIL_SCREEN",
		Render: func(m exploreModel) string {
			selectedUser := m.users[m.selectedUserIndex]
			jsonString, err := json.MarshalIndent(selectedUser.Claims, "", "  ")
			if err != nil {
				panic(err)
			}
			m.textArea.SetValue(string(jsonString))
			return m.textArea.View()
		}}
)

type exploreModel struct {
	selectedScreen Screen

	users             []authentication.User
	selectedUserIndex int

	searchQuery     string
	filteredResults []authentication.User

	textArea textarea.Model
}

func createExploreModel(users []authentication.User) exploreModel {

	return exploreModel{
		selectedScreen: LIST_SCREEN,

		users:             users,
		selectedUserIndex: 0,

		searchQuery:     "",
		filteredResults: []authentication.User{},

		textArea: textarea.New(),
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
			if m.selectedScreen.Name == LIST_SCREEN.Name {
				clearScreen()
				return m, tea.Quit
			} else {
				m.selectedScreen = LIST_SCREEN
				break
			}
		case "j", "down":
			m.MoveCursorDown()
			break
		case "k", "up":
			m.MoveCursorUp()
			break
		case "enter":
			m.ViewDetail()
			break
		}
	}

	return m, nil
}

func (m *exploreModel) ViewDetail() {
	m.selectedScreen = DETAIL_SCREEN
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
	output := m.selectedScreen.Render(m)
	return output
}
