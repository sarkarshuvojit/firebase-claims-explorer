package components

import (
	"fmt"

	"shuvojit.in/firebase-claims-explorer/authentication"
)

func UserList(users []authentication.User, selectedUser authentication.User) string {
	output := ""
	for _, user := range users {
		if user.UID == selectedUser.UID {
			output += fmt.Sprintf("> %s - %s\n", user.UID, user.Email)
		} else {
			output += fmt.Sprintf("%s - %s\n", user.UID, user.Email)
		}
	}
	return output
}
