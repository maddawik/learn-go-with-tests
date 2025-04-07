package poker_test

import (
	"fmt"
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
		dummySpyAlerter := &SpyBlindAlerter{}

		cli := poker.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "May")
	})

	t.Run("record Cody win from user input", func(t *testing.T) {
		in := strings.NewReader("Cody wins\n")
		playerStore := &poker.StubPlayerStore{}
		dummySpyAlerter := &SpyBlindAlerter{}

		cli := poker.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cody")
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("James wins\n")
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := poker.NewCLI(playerStore, in, blindAlerter)
		cli.PlayPoker()

		cases := []struct {
			expectedScheduleTime time.Duration
			expectedAmount       int
		}{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}

		for i, c := range cases {
			t.Run(fmt.Sprintf("%d schedule for %v", c.expectedAmount, c.expectedScheduleTime), func(t *testing.T) {
				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
				}

				alert := blindAlerter.alerts[i]

				amountGot := alert.amount
				if amountGot != c.expectedAmount {
					t.Errorf("got amount %d, want %d", amountGot, c.expectedAmount)
				}

				gotScheduledTime := alert.scheduleAt
				if gotScheduledTime != c.expectedScheduleTime {
					t.Errorf("got time %v, want %v", gotScheduledTime, c.expectedScheduleTime)
				}
			})
		}
	})
}
