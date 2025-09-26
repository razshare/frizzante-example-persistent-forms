package welcome

import (
	"main/lib/core/client"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/core/view"
	session "main/lib/session/memory"
	"os"
)

func View(client *client.Client) {
	send.FileOrElse(client, send.FileOrElseConfig{
		UseDisk: os.Getenv("DEV") == "1",
		OrElse: func() {
			state := session.Start(receive.SessionId(client))
			send.View(client, view.View{
				Name:  "Welcome",
				Props: state.Form,
			})
		},
	})
}
