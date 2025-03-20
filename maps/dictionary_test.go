package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "testing things"}

	got := Search(dictionary, "test")
	want := "testing things"

	if got != want {
		t.Errorf("got %q, want %q given %q", got, want, dictionary)
	}
}
