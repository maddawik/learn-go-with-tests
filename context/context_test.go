package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type SpyStore struct {
	response string
}

func (s *SpyStore) Fetch() string {
	return s.response
}

func TestServer(t *testing.T) {
	data := "hello, world"
	svr := Server(&SpyStore{data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	svr.ServeHTTP(response, request)

	if response.Body.String() != data {
		t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
	}
}
