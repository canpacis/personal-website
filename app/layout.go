package app

import (
	"github.com/canpacis/pacis/pages"
	font "github.com/canpacis/pacis/pages/font"
	. "github.com/canpacis/pacis/ui/html"
)

var inter = font.New("Inter", font.WeightList{font.W100, font.W900}, font.Swap, font.Latin, font.LatinExt)

func Layout(ctx *pages.LayoutContext) I {
	return Html(
		Lang("en"),
		Class("dark"),

		Head(
			font.Head(inter),
			ctx.Head(),
			Link(Href("/public/main.css"), Rel("stylesheet")),
			Script(Src("/public/main.js")),
			Link(Href("/public/profile.webp"), Rel("icon"), Type("image/webp")),
			Title(Text("canpacis | Muhammed Ali CAN")),
		),
		Body(
			Class("flex flex-col min-h-screen"),

			ctx.Outlet(),
		),
	)
}
