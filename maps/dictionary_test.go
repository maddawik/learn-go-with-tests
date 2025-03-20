package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "testing things"}

	got := Search(dictionary, "test")
	want := "testing things"

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
