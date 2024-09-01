package handlers_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	"github.com/subterraneanbob/go-complete-web/internal/dict"
	"github.com/subterraneanbob/go-complete-web/internal/handlers"
)

func TestSearchHandlerHtml(t *testing.T) {
	data := search(t, "text/html")
	approvals.VerifyString(t, data)
}

func TestSearchHandlerJson(t *testing.T) {
	data := search(t, "application/json")
	approvals.VerifyString(t, data)
}

func search(t *testing.T, accept string) string {

	req := httptest.NewRequest(http.MethodGet, "/search?searchTerm=table", nil)
	req.Header.Set("Accept", accept)

	w := httptest.NewRecorder()
	handlers.HandleSearch(dict.NewDict()).ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	return string(data)
}
