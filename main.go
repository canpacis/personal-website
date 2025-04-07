package main

import (
	"embed"
	"net/http"
	"os"

	"github.com/canpacis/pacis/pages"
	"github.com/canpacis/personal-website/app"
)

//go:embed public
var public embed.FS

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
	router := pages.Routes(
		pages.Public(public, "public"),
		pages.Layout(app.Layout),
		pages.Route(pages.Path("/"), pages.Page(app.HomePage)),
	)

	mux := router.Handler().(*http.ServeMux)
	mux.Handle("GET /robots.txt", fileServer(robots, "text/plain; charset=utf-8"))
	mux.Handle("GET /sitemap.xml", fileServer(sitemap, "application/xml"))

	http.ListenAndServe(":"+getEnv("PORT", "8080"), mux)
}
