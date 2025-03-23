package concurrency

import "net/http"

func CheckWebsite(url string) bool {
	r, err := http.Head(url)
	if err != nil {
		return false
	}

	return r.StatusCode == http.StatusOK
}
