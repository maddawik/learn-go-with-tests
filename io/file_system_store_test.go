package main

import (
	"io"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Fox", "Wins": 10},
			{"Name": "Falco", "Wins": 33}]`)
		defer cleanDatabase()

		store := FileSystemPlayerStore{database}

		want := []Player{
			{Name: "Fox", Wins: 10},
			{Name: "Falco", Wins: 33},
		}

		got := store.GetLeague()
		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Fox", "Wins": 10},
			{"Name": "Falco", "Wins": 33}]`)
		defer cleanDatabase()

		store := FileSystemPlayerStore{database}

		got := store.GetPlayerScore("Fox")
		want := 10

		assertPlayerScore(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Jimmy", "Wins": 2},
			{"Name": "Jimbo", "Wins": 4}]`)
		defer cleanDatabase()

		store := FileSystemPlayerStore{database}

		player := "Jimbo"

		store.RecordWin(player)

		got := store.GetPlayerScore(player)
		want := 5

		assertPlayerScore(t, got, want)
	})
}

func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpfile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

func assertPlayerScore(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
