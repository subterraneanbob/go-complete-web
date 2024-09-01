package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/subterraneanbob/go-complete-web/internal/dict"
	"github.com/subterraneanbob/go-complete-web/internal/handlers"
	static "github.com/subterraneanbob/go-complete-web/web"
)

func main() {

	d := dict.NewMapDict()

	mux := http.NewServeMux()

	mux.Handle("/dist/", NoDirView(http.FileServerFS(static.DistFS)))
	mux.Handle("/search", LogRequest(handlers.HandleSearch(d)))
	mux.Handle("/", LogRequest(handlers.HandleIndex(d)))

	addr := "localhost:8080"
	log.Printf("Listening on %v", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}

// NoDirView removes the ability to view the folder listing from http.FileServerFS.
// It responds with 404 Not Found when no specific file is requested.
func NoDirView(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func LogRequest(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v %v", r.Method, r.URL)
		h.ServeHTTP(w, r)
	})
}
