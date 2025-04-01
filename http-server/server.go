package main

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	if player == "Fox" {
		fmt.Fprintf(w, "20")
	}
	if player == "Falco" {
		fmt.Fprintf(w, "10")
	}
}
