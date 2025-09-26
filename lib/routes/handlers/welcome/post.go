package welcome

import (
	"main/lib/core/client"
	"main/lib/core/receive"
	"main/lib/core/send"
	session "main/lib/session/memory"
)

func Post(client *client.Client) {
	var errorMessage string

	state := session.Start(receive.SessionId(client))
	username := receive.FormValue(client, "username")

	if len(username) < 3 {
		// Explain why the input is invalid.
		errorMessage = "username is too short"
	} else {
		// If the input is valid, reset it.
		username = ""
	}

	state.Form = session.Form{
		Username: username,
		Error:    errorMessage,
	}

	send.Navigate(client, "/")
}
