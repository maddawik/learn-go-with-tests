package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Wendell")

	got := buffer.String()
	want := "Hello, Wendell"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
