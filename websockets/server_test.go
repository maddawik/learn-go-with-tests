package poker

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

var dummyGame = &GameSpy{}

func TestGETPlayers(t *testing.T) {
	store := &StubPlayerStore{
		scores: map[string]int{
			"Fox":   20,
			"Falco": 10,
		},
		winCalls: nil,
	}
	server := mustMakePlayerServer(t, store, dummyGame)
	t.Run("returns Fox's score", func(t *testing.T) {
		request := NewGetScoreRequest("Fox")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Falco's score", func(t *testing.T) {
		request := NewGetScoreRequest("Falco")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on player that doesn't exist", func(t *testing.T) {
		request := NewGetScoreRequest("Slippy")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := &StubPlayerStore{
		scores:   map[string]int{},
		winCalls: []string{},
	}
	server := mustMakePlayerServer(t, store, dummyGame)
	t.Run("it returns accepted on POST", func(t *testing.T) {
		player := "Fox"

		request := NewPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response, http.StatusAccepted)
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
		server := mustMakePlayerServer(t, store, dummyGame)

		request := NewLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := GetLeagueFromResponse(t, response.Body)

		AssertContentType(t, response, "application/json")
		assertStatus(t, response, http.StatusOK)
		AssertLeague(t, got, wantedLeague)
	})
}

func TestGame(t *testing.T) {
	t.Run("GET /game returns 200", func(t *testing.T) {
		server := mustMakePlayerServer(t, &StubPlayerStore{}, dummyGame)

		request := NewGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response, http.StatusOK)
	})
	t.Run("start a game with 3 players, send some blind alerts down the WS and declare Jimbo the winner", func(t *testing.T) {
		wantedBlindAlert := "Blind is 100"
		winner := "Jimbo"

		game := &GameSpy{BlindAlert: []byte(wantedBlindAlert)}
		dummyPlayerStore := &StubPlayerStore{}

		server := httptest.NewServer(mustMakePlayerServer(t, dummyPlayerStore, game))
		ws := mustDialWS(t, "ws"+strings.TrimPrefix(server.URL, "http")+"/ws")

		defer server.Close()
		defer ws.Close()

		writeWSMessage(t, ws, "3")
		writeWSMessage(t, ws, winner)

		waitTime := 10 * time.Millisecond
		time.Sleep(waitTime)
		AssertGameStartedWith(t, 3, game)
		AssertGameFinishedWith(t, winner, game)

		within(t, waitTime, func() { assertWebsocketGotMsg(t, ws, wantedBlindAlert) })
	})
}

func assertWebsocketGotMsg(t *testing.T, ws *websocket.Conn, want string) {
	t.Helper()
	_, got, _ := ws.ReadMessage()

	if string(got) != want {
		t.Errorf("got blind alert %q, wanted %q", string(got), want)
	}
}

func within(t testing.TB, d time.Duration, assert func()) {
	t.Helper()

	done := make(chan struct{}, 1)

	go func() {
		assert()
		done <- struct{}{}
	}()

	select {
	case <-time.After(d):
		t.Error("timed out")
	case <-done:
	}
}

func writeWSMessage(t *testing.T, ws *websocket.Conn, message string) {
	err := ws.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		t.Fatalf("could not second message over ws connection %v", err)
	}
}

func mustDialWS(t *testing.T, url string) *websocket.Conn {
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("could not open a ws conntection on %s %v", url, err)
	}
	return ws
}

func assertStatus(t testing.TB, got *httptest.ResponseRecorder, want int) {
	t.Helper()
	if got.Code != want {
		t.Errorf("didn't get correct status, got %d want %d", got.Code, want)
	}
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

func mustMakePlayerServer(t *testing.T, store PlayerStore, game Game) *PlayerServer {
	server, err := NewPlayerServer(store, game)
	if err != nil {
		t.Fatal("problem creating player server", err)
	}
	return server
}
