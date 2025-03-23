package racer

import "testing"

func TestRacer(t *testing.T) {
	fastUrl := "https://ejreilly.xyz"
	slowUrl := "https://google.com"

	result := Racer(fastUrl, slowUrl)

	if result != fastUrl {
		t.Fatalf("expected %v, but got %v", fastUrl, result)
	}
}
