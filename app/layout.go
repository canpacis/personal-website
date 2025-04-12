package app

import (
	"github.com/canpacis/pacis/pages"
	font "github.com/canpacis/pacis/pages/font"
	"github.com/canpacis/pacis/pages/i18n"
	. "github.com/canpacis/pacis/ui/html"
	"golang.org/x/text/language"
)

var inter = font.New("Inter", font.WeightList{font.W100, font.W900}, font.Swap, font.Latin, font.LatinExt)

func Layout(ctx *pages.LayoutContext) I {
	locale, err := i18n.Locale(ctx)
	if err != nil {
		locale = &language.English
	}
	title := i18n.Text("title").String(ctx)
	desc := i18n.Text("subtitle").String(ctx)
	keywords := "canpacis, pacis, ui/ux, frontend developer, library author, Muhammed Ali CAN, Go, JavaScript, TypeScript, Next JS, React, ui library"

	return Html(
		Lang(locale.String()),
		Class("dark"),

		Head(
			font.Head(inter),
			ctx.Head(),

			Meta(Name("title"), Content(title)),
			Meta(Name("description"), Content(desc)),
			Meta(Name("keywords"), Content(keywords)),
			Meta(Name("robots"), Content("index, follow")),
			Meta(HttpEquiv("Content-Type"), Content("text/html; charset=utf-8")),
			Meta(HttpEquiv("language"), Content("English")),
			Meta(HttpEquiv("author"), Content("canpacis")),

			Meta(Property("og:type"), Content("website")),
			Meta(Property("og:url"), Content("https://ui.canpacis.com")),
			Meta(Property("og:title"), Content(title)),
			Meta(Property("og:description"), Content(desc)),
			Meta(Property("og:image"), Content("/public/banner.webp")),

			Meta(Property("twitter:card"), Content("summary_large_image")),
			Meta(Property("twitter:url"), Content("https://ui.canpacis.com")),
			Meta(Property("twitter:title"), Content(title)),
			Meta(Property("twitter:description"), Content(desc)),
			Meta(Property("twitter:image"), Content("/public/banner.webp")),

			Title(Text(title)),

			Link(Href("/public/main.css"), Rel("stylesheet")),
			Script(Src("/public/main.js")),
			Link(Href("/public/profile.webp"), Rel("icon"), Type("image/webp")),
		),
		Body(
			Class("flex flex-col min-h-screen"),

			ctx.Outlet(),
		),
	)
}
