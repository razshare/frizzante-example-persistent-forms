//go:build !types

package main

import (
	"embed"
	"os"

	"main/lib/core/route"
	"main/lib/core/server"
	"main/lib/core/view/ssr"
	"main/lib/routes/handlers/welcome"
)

//go:generate make clean configure
//go:generate make package
//go:generate make types
//go:embed app/dist
var efs embed.FS
var srv = server.New()
var dev = os.Getenv("DEV") == "1"
var render = ssr.New(ssr.Config{Efs: efs, UseDisk: dev})

func main() {
	defer server.Start(srv)
	srv.Efs = efs
	srv.Render = render
	srv.Routes = []route.Route{
		{Pattern: "GET /", Handler: welcome.View},
		{Pattern: "POST /", Handler: welcome.Submit},
	}
}
