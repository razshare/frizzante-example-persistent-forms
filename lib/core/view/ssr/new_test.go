package ssr

import (
	"embed"
	"strings"
	"testing"

	"main/lib/core/view"
)

//go:generate rm -fr ./app
//go:generate mkdir -p ./app
//go:generate cp -r ../../../../app/dist ./app
//go:embed app
var TestNewEfs embed.FS

func TestNew(t *testing.T) {
	render := New(Config{Efs: TestNewEfs})
	html, err := render(view.View{Name: "Welcome"})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(html, "Show Todos") {
		t.Fatal("view should contain Show Todos")
	}
}
