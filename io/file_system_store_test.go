package main

import (
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Fox", "Wins": 10},
			{"Name": "Falco", "Wins": 33}]`)

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
		database := strings.NewReader(`[
			{"Name": "Fox", "Wins": 10},
			{"Name": "Falco", "Wins": 33}]`)

		store := FileSystemPlayerStore{database}

		got := store.GetPlayerScore("Fox")
		want := 10

		assertPlayerScore(t, got, want)
	})
}

func assertPlayerScore(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
