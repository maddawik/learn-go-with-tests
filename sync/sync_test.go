package main

import "testing"

func TestCounter(t *testing.T) {
	t.Run("increment the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		got := counter.Value()
		want := 3

		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	})
}
