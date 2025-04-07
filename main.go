package main

import (
	"embed"
	"net/http"

	"github.com/canpacis/pacis/pages"
	"github.com/canpacis/personal-website/app"
)

//go:embed public
var public embed.FS

func main() {
	router := pages.Routes(
		pages.Public(public, "public"),
		pages.Layout(app.Layout),
		pages.Route(pages.Path("/"), pages.Page(app.HomePage)),
	)

	http.ListenAndServe(":8081", router.Handler())
}
