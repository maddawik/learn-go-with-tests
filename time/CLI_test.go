package poker_test

import (
	"strings"
	"testing"
	"time"

	poker "github.com/maddawik/learn-go-with-tests/time"
)

type SpyBlindAlerter struct {
	alerts []struct {
		scheduleAt time.Duration
		amount     int
	}
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, struct {
		scheduleAt time.Duration
		amount     int
	}{duration, amount})
}

func TestCLI(t *testing.T) {
	t.Run("record May win from user input", func(t *testing.T) {
		in := strings.NewReader("May wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "May")
	})

	t.Run("record Cody win from user input", func(t *testing.T) {
		in := strings.NewReader("Cody wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cody")
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("James wins\n")
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := poker.NewCLI(playerStore, in, blindAlerter)
		cli.PlayPoker()

		if len(blindAlerter.alerts) != 1 {
			t.Fatal("expected a blind alert to be scheduled")
		}
	})
}
