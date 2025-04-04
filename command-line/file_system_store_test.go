package poker

import (
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name": "Fox", "Wins": 10},
			{"Name": "Falco", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		AssertNoError(t, err)

		want := []Player{
			{Name: "Falco", Wins: 33},
			{Name: "Fox", Wins: 10},
		}

		got := store.GetLeague()
		AssertLeague(t, got, want)

		got = store.GetLeague()
		AssertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name": "Fox", "Wins": 10},
			{"Name": "Falco", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		AssertNoError(t, err)

		got := store.GetPlayerScore("Fox")
		want := 10

		AssertPlayerScore(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name": "Jimmy", "Wins": 2},
			{"Name": "Jimbo", "Wins": 4}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		AssertNoError(t, err)

		player := "Jimbo"

		store.RecordWin(player)

		got := store.GetPlayerScore(player)
		want := 5

		AssertPlayerScore(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name": "Cloud", "Wins": 8},
			{"Name": "Barret", "Wins": 5}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		AssertNoError(t, err)

		player := "Tifa"

		store.RecordWin(player)

		got := store.GetPlayerScore(player)
		want := 1
		AssertPlayerScore(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPlayerStore(database)

		AssertNoError(t, err)
	})

	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name": "Cody", "Wins": 2},
			{"Name": "May", "Wins": 10}]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		AssertNoError(t, err)

		want := []Player{
			{Name: "May", Wins: 10},
			{Name: "Cody", Wins: 2},
		}

		got := store.GetLeague()
		AssertLeague(t, got, want)

		// check again
		got = store.GetLeague()
		AssertLeague(t, got, want)
	})
}

func CreateTempFile(t testing.TB, initialData string) (*os.File, func()) {
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

func AssertPlayerScore(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func AssertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("didn't expect an error but got one, %v", err)
	}
}
