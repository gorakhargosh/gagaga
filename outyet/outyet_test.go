package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestIntegration(t *testing.T) {
	status := statusHandler(404)
	ts := httptest.NewServer(&status)
	defer ts.Close()

	sleep := make(chan bool)
	pollSleep = func(time.Duration) {
		sleep <- true
		sleep <- true
	}
	done := make(chan bool)
	pollDone = func() {
		done <- true
	}

	s := NewServer("1.x", ts.URL, 1*time.Millisecond)
	<-sleep
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	s.ServeHTTP(w, r)
	if b := w.Body.String(); !strings.Contains(b, "No.") {
		t.Errorf("body = %q, wanted no", b)
	}

	status = 200
	<-sleep
	<-done
	w = httptest.NewRecorder()
	s.ServeHTTP(w, r)
	if b := w.Body.String(); !strings.Contains(b, "YES!") {
		t.Errorf("body = %q, wanted yes!", b)
	}
}

type statusHandler int

func (h *statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(int(*h))
}

func TestIsTagged(t *testing.T) {
	status := statusHandler(404)
	ts := httptest.NewServer(&status)
	defer ts.Close()

	if isTagged(ts.URL) {
		t.Error("isTagged returned true, want false")
	}

	status = 200
	if !isTagged(ts.URL) {
		t.Error("isTagged returned false, want true")
	}
}
