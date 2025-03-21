package main

import "testing"

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "metroid"
		definition := "prime"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "paper mario"
		definition := "the origami king"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "the thousand-year door")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
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

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "paper mario"
		definition := "the origami king"
		dictionary := Dictionary{word: definition}
		newDefinition := "the thousand-year door"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "metroid"
		definition := "fusion"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
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
