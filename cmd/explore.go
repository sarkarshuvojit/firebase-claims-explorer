package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"shuvojit.in/firebase-claims-explorer/authentication"
	"shuvojit.in/firebase-claims-explorer/components"
)

var (

	// General.

	subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}

	divider = lipgloss.NewStyle().
		SetString("•").
		Padding(0, 1).
		Foreground(subtle).
		String()

	url = lipgloss.NewStyle().Foreground(special).Render

	// Tabs.

	activeTabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      " ",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┘",
		BottomRight: "└",
	}

	tabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┴",
		BottomRight: "┴",
	}

	tab = lipgloss.NewStyle().
		Border(tabBorder, true).
		BorderForeground(highlight).
		Padding(0, 1)

	activeTab = tab.Copy().Border(activeTabBorder, true)

	tabGap = tab.Copy().
		BorderTop(false).
		BorderLeft(false).
		BorderRight(false)

	// Title.

	titleStyle = lipgloss.NewStyle().
			MarginLeft(1).
			MarginRight(5).
			Padding(0, 1).
			Italic(true).
			Foreground(lipgloss.Color("#FFF7DB")).
			SetString("Lip Gloss")

	descStyle = lipgloss.NewStyle().MarginTop(1)

	infoStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderTop(true).
			BorderForeground(subtle)

	// Dialog.

	dialogBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			Padding(1, 0).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true)

	buttonStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFF7DB")).
			Background(lipgloss.Color("#888B7E")).
			Padding(0, 3).
			MarginTop(1)

	activeButtonStyle = buttonStyle.Copy().
				Foreground(lipgloss.Color("#FFF7DB")).
				Background(lipgloss.Color("#F25D94")).
				MarginRight(2).
				Underline(true)

	// List.

	list = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, true, false, false).
		BorderForeground(subtle).
		MarginRight(2).
		Height(8).
		Width(columnWidth + 1)

	listHeader = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			BorderForeground(subtle).
			MarginRight(2).
			Render

	listItem = lipgloss.NewStyle().PaddingLeft(2).Render

	checkMark = lipgloss.NewStyle().SetString("✓").
			Foreground(special).
			PaddingRight(1).
			String()

	listDone = func(s string) string {
		return checkMark + lipgloss.NewStyle().
			Strikethrough(true).
			Foreground(lipgloss.AdaptiveColor{Light: "#969B86", Dark: "#696969"}).
			Render(s)
	}

	// Paragraphs/History.

	historyStyle = lipgloss.NewStyle().
			Align(lipgloss.Left).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(highlight).
			Margin(1, 3, 0, 0).
			Padding(1, 2).
			Height(19).
			Width(columnWidth)

	// Status Bar.

	statusNugget = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Padding(0, 1)

	statusBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
			Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})

	statusStyle = lipgloss.NewStyle().
			Inherit(statusBarStyle).
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#FF5F87")).
			Padding(0, 1).
			MarginRight(1)

	encodingStyle = statusNugget.Copy().
			Background(lipgloss.Color("#A550DF")).
			Align(lipgloss.Right)

	statusText = lipgloss.NewStyle().Inherit(statusBarStyle)

	fishCakeStyle = statusNugget.Copy().Background(lipgloss.Color("#6124DF"))

	borderStyle = lipgloss.NewStyle().
			Border(lipgloss.DoubleBorder()).
			Render

	// Page.

	docStyle = lipgloss.NewStyle().Align(lipgloss.Center)
)

const (
	// In real life situations we'd adjust the document to fit the width we've
	// detected. In the case of this example we're hardcoding the width, and
	// later using the detected width only to truncate in order to avoid jaggy
	// wrapping.
	width = 96

	columnWidth = 30
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
			output := components.UserList(m.users, m.users[m.selectedUserIndex])
			return output
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
			m.textArea.SetHeight(lipgloss.Height(m.textArea.Value()))
			style := lipgloss.NewStyle().
				MarginLeft(20)
			return style.Render(m.textArea.View())
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
			clearScreen()
			return m, tea.Quit
		case "j", "down":
			m.MoveCursorDown()
			break
		case "k", "up":
			m.MoveCursorUp()
			break
		case "enter":
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

func withTitle(title string, body string) string {
	return lipgloss.JoinVertical(
		lipgloss.Center,
		listHeader(title),
		body,
	)
}

func (m exploreModel) View() string {
	physicalWidth, physicalHeight, _ := term.GetSize(int(os.Stdout.Fd()))
	w := lipgloss.Width
	h := lipgloss.Height
	doc := strings.Builder{}
	{
		doc.WriteString(
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				withTitle("Users", LIST_SCREEN.Render(m)),
				withTitle("Claims", DETAIL_SCREEN.Render(m)),
			))
		doc.WriteString("\n\n")
	}
	{

		statusKey := statusStyle.Render("STATUS")
		statusVal := statusText.Copy().
			Width(width - w(statusKey)).
			Render("⏶/k: up | ⏷/j: down | q: quit")

		bar := lipgloss.JoinHorizontal(lipgloss.Top,
			statusKey,
			statusVal,
		)

		diff := physicalHeight - h(doc.String())
		marginTop := lipgloss.NewStyle().
			MarginTop(diff).
			Render

		doc.WriteString(marginTop(statusBarStyle.Width(physicalWidth).Render(bar)))

	}

	if physicalWidth > 0 {
		docStyle = docStyle.MaxWidth(physicalWidth)
	}

	return docStyle.Render(doc.String())
}
