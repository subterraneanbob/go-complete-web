package handlers_test

import (
	"os"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestMain(m *testing.M) {
	approvals.UseFolder("testdata")
	os.Exit(m.Run())
}
