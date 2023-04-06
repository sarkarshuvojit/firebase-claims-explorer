package components

import (
	"github.com/charmbracelet/lipgloss"
	"shuvojit.in/firebase-claims-explorer/authentication"
)

var (
	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
			Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"}).
			Render
	defaultStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
			Render
)

func UserList(users []authentication.User, selectedUser authentication.User) string {
	var output []string
	for _, user := range users {
		if user.UID == selectedUser.UID {
			output = append(output, selectedStyle(user.UID))
		} else {
			output = append(output, defaultStyle(user.UID))
		}
	}
	return lipgloss.JoinVertical(
		lipgloss.Right,
		output...,
	)
}
