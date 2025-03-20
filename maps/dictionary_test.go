package main

import "testing"

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "testing things"}

		got, _ := dictionary.Search("test")
		want := "testing things"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "testing things"}

		_, err := dictionary.Search("metroid")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertStrings(t, err.Error(), ErrNotFound.Error())
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
