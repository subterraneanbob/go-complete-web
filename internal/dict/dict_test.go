package dict_test

import (
	"slices"
	"testing"

	"github.com/subterraneanbob/go-complete-web/internal/dict"
)

func TestDictFind(t *testing.T) {

	testAllFindersWith(t, func(f dict.Finder) {
		term := "cherryblossom"
		got := f.Find(term, 5)
		want := []string{term}

		if !slices.Equal(got, want) {
			t.Errorf("Find() did not return correct results: got %v, want %v", got, want)
		}
	})
}

func TestDictFindLimit(t *testing.T) {

	testAllFindersWith(t, func(f dict.Finder) {
		term := "a"
		got := f.Find(term, 10)
		want := []string{"a", "aa", "aaa", "aah", "aahed", "aahing", "aahs", "aal", "aalii", "aaliis"}

		if !slices.Equal(got, want) {
			t.Errorf("Find() did not return correct results: got %v, want %v", got, want)
		}
	})
}

func BenchmarkDictNew(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_ = dict.NewDict()
	}
}

func BenchmarkDictFind(b *testing.B) {

	d := dict.NewDict()

	for i := 0; i < b.N; i++ {
		d.Find("ubiquious", 5)
	}
}

func BenchmarkMapDictFind(b *testing.B) {

	d := dict.NewMapDict()

	for i := 0; i < b.N; i++ {
		d.Find("ubiquious", 5)
	}
}

func BenchmarkMapDictNew(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_ = dict.NewMapDict()
	}
}

func testAllFindersWith(t *testing.T, testFunc func(f dict.Finder)) {

	t.Helper()

	cases := []struct {
		name   string
		finder dict.Finder
	}{
		{
			"dict",
			dict.NewDict(),
		},
		{
			"dict_map",
			dict.NewMapDict(),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			testFunc(c.finder)
		})
	}
}
