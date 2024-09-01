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

func TestIndexHandler(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/?q=monkey", nil)
	w := httptest.NewRecorder()
	d := dict.NewDict()

	handlers.HandleIndex(d).ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	approvals.VerifyString(t, string(data))
}
