package main

import "testing"

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	word := "metroid"
	definition := "prime"

	dictionary.Add(word, definition)

	assertDefinition(t, dictionary, word, definition)
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "testing things"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "testing things"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("metroid")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrNotFound)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("shouldn't have got an error:", err)
	}

	assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
