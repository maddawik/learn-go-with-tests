package poker

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := CreateTempFile(t, "[]")
	defer cleanDatabase()
	store, err := NewFileSystemPlayerStore(database)

	AssertNoError(t, err)

	server := NewPlayerServer(store)
	player := "Falco"

	server.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, NewGetScoreRequest(player))

		AssertStatus(t, response.Code, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, NewLeagueRequest())

		got := GetLeagueFromResponse(t, response.Body)
		want := []Player{
			{Name: player, Wins: 3},
		}

		AssertStatus(t, response.Code, http.StatusOK)
		AssertLeague(t, got, want)
	})
}
