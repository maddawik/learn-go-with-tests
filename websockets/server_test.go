package poker

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	store := &StubPlayerStore{
		scores: map[string]int{
			"Fox":   20,
			"Falco": 10,
		},
		winCalls: nil,
	}
	server := NewPlayerServer(store)
	t.Run("returns Fox's score", func(t *testing.T) {
		request := NewGetScoreRequest("Fox")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Falco's score", func(t *testing.T) {
		request := NewGetScoreRequest("Falco")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on player that doesn't exist", func(t *testing.T) {
		request := NewGetScoreRequest("Slippy")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := &StubPlayerStore{
		scores:   map[string]int{},
		winCalls: []string{},
	}
	server := NewPlayerServer(store)
	t.Run("it returns accepted on POST", func(t *testing.T) {
		player := "Fox"

		request := NewPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusAccepted)
		AssertPlayerWin(t, store, player)
	})
}

func TestLeague(t *testing.T) {
	t.Run("it returns 200 on /league", func(t *testing.T) {
		wantedLeague := []Player{
			{Name: "EJ", Wins: 2},
			{Name: "Kayla", Wins: 10},
		}

		store := &StubPlayerStore{league: wantedLeague}
		server := NewPlayerServer(store)

		request := NewLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := GetLeagueFromResponse(t, response.Body)

		AssertContentType(t, response, "application/json")
		AssertStatus(t, response.Code, http.StatusOK)
		AssertLeague(t, got, wantedLeague)
	})
}

func TestGame(t *testing.T) {
	t.Run("GET /game returns 200", func(t *testing.T) {
		server := NewPlayerServer(&StubPlayerStore{})

		request := NewGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		AssertStatus(t, response.Code, http.StatusOK)
	})
}

func NewGameRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/game", nil)
	return request
}

func NewLeagueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}

func NewGetScoreRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func NewPostWinRequest(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func AssertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body incorrect, got %q want %q", got, want)
	}
}

func AssertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response didn't have content-type of %s, %v", response.Result().Header, want)
	}
}

func AssertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("didn't get correct status, got %d want %d", got, want)
	}
}

func AssertLeague(t testing.TB, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func GetLeagueFromResponse(t testing.TB, body io.Reader) []Player {
	t.Helper()

	league, err := NewLeague(body)
	if err != nil {
		t.Fatalf("unable to parson json response from server %q into slice of Player, '%v'", body, err)
	}

	return league
}
