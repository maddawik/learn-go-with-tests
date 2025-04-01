package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (s *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &PlayerServer{store: &InMemoryPlayerStore{}}
	handler := http.HandlerFunc(server.ServeHTTP)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
