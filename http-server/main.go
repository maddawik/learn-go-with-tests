package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func (s *InMemoryPlayerStore) RecordWin(name string) {
}

func main() {
	server := &PlayerServer{store: &InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}
