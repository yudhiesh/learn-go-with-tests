package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayedServer(duration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(duration)
		w.WriteHeader(http.StatusOK)
	}))

}

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()
		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("returns an error if there is a timeout", func(t *testing.T) {
		serverA := makeDelayedServer(20 * time.Millisecond)
		serverB := makeDelayedServer(25 * time.Millisecond)
		defer serverA.Close()
		defer serverB.Close()
		slowURL := serverA.URL
		fastURL := serverB.URL

		_, err := ConfigurableRacer(slowURL, fastURL, time.Millisecond*10)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}
