package welcome

import (
	"main/lib/core/client"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/core/view"
	session "main/lib/session/memory"
)

func View(client *client.Client) {
	state := session.Start(receive.SessionId(client))
	send.View(client, view.View{
		Name:  "Welcome",
		Props: state.Form,
	})
}
