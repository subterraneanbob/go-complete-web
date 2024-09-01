package handlers

import (
	"net/http"

	"github.com/subterraneanbob/go-complete-web/internal/dict"
	"github.com/subterraneanbob/go-complete-web/web/templates"
)

func HandleIndex(d dict.Finder) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var matches []string

		searchTerm := r.URL.Query().Get("q")
		if searchTerm != "" {
			matches = d.Find(searchTerm, wordsLimit)
		}

		component := templates.Index(searchTerm, matches)
		component.Render(r.Context(), w)
	})
}
