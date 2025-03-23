package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	return url == "https://metroid.prime"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://ejreilly.xyz",
		"https://metroid.prime",
	}

	want := map[string]bool{
		"https://google.com":    false,
		"https://ejreilly.xyz":  false,
		"https://metroid.prime": true,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
