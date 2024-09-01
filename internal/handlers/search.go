package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/subterraneanbob/go-complete-web/internal/dict"
	"github.com/subterraneanbob/go-complete-web/web/templates"
)

const wordsLimit = 10

func HandleSearch(d dict.Finder) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		searchTerm := r.URL.Query().Get("searchTerm")
		matches := d.Find(searchTerm, wordsLimit)

		accept := r.Header.Get("Accept")
		if strings.Contains(accept, "application/json") {

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(matches)

		} else {

			w.Header().Set("Content-Type", "text/html; charset=utf-8")

			component := templates.Results(searchTerm, matches)
			component.Render(r.Context(), w)
		}
	})
}
