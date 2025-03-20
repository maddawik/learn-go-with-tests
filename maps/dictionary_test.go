package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "testing things"}

	got := dictionary.Search("test")
	want := "testing things"

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
