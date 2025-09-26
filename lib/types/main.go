package main

import (
	"main/lib/core/types"
	session "main/lib/session/memory"
)

func main() {
	types.Generate[session.Form]()
}
