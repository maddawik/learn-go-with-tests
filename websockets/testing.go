package poker

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

func AssertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
	}

	if store.winCalls[0] != winner {
		t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], winner)
	}
}

type SpyBlindAlerter struct {
	Alerts []ScheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int, to io.Writer) {
	s.Alerts = append(s.Alerts, ScheduledAlert{duration, amount})
}

type ScheduledAlert struct {
	At     time.Duration
	Amount int
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}

type GameSpy struct {
	StartCalled bool
	StartedWith int

	FinishedWith   string
	FinishedCalled bool
}

func (g *GameSpy) Play(numberOfPlayers int, to io.Writer) {
	g.StartCalled = true
	g.StartedWith = numberOfPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedCalled = true
	g.FinishedWith = winner
}

func AssertGameNotStarted(t testing.TB, started bool) {
	t.Helper()
	if started {
		t.Error("game should not have started")
	}
}

func AssertGameStarted(t testing.TB, started bool) {
	t.Helper()
	if !started {
		t.Error("game should have started")
	}
}

func AssertGameStartedWith(t testing.TB, want, got int) {
	t.Helper()
	if got != want {
		t.Errorf("game should have started with %q but got %q", want, got)
	}
}

func AssertGameFinishedWith(t testing.TB, want, got string) {
	t.Helper()
	if got != want {
		t.Errorf("game should have finished with %q but got %q", want, got)
	}
}

func AssertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
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
