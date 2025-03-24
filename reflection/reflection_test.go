package main

import "testing"

func TestWalk(t *testing.T) {
	want := "Kayla"
	var got []string

	x := struct {
		Name string
	}{want}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d wanted %d", len(got), 1)
	}
}
