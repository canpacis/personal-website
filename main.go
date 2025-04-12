package main

import (
	"embed"
	"log"
	"net/http"
	"os"

	"github.com/canpacis/pacis/pages"
	"github.com/canpacis/pacis/pages/i18n"
	"github.com/canpacis/pacis/pages/middleware"
	"github.com/canpacis/personal-website/app"
	"golang.org/x/text/language"
)

//go:embed public
var public embed.FS

//go:embed messages
var messages embed.FS

//go:embed robots.txt
var robots []byte

//go:embed sitemap.xml
var sitemap []byte

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func fileServer(data []byte, contenttyp string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", contenttyp)
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})
}

func main() {
	bundle, err := i18n.Setup(messages, language.English)
	if err != nil {
		log.Fatal(err)
	}

	router := pages.Routes(
		pages.Public(public, "public"),
		pages.Middleware(middleware.Locale(bundle, language.English)),
		pages.Layout(app.Layout),
		pages.Route(pages.Path("/"), pages.Page(app.HomePage)),
	)

	mux := router.Handler().(*http.ServeMux)
	mux.Handle("GET /robots.txt", fileServer(robots, "text/plain; charset=utf-8"))
	mux.Handle("GET /sitemap.xml", fileServer(sitemap, "application/xml"))

	http.ListenAndServe(":"+getEnv("PORT", "8080"), mux)
}
