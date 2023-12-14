package main

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newTestApp() *application {
	return &application{
		logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}
}

type testServer struct {
	*httptest.Server
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewServer(h)
	return &testServer{ts}
}

func (ts *testServer) get(t *testing.T, url string) (int, http.Header, string) {
	rs, err := ts.Client().PostForm(ts.URL+url, nil)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace(body)
	return rs.StatusCode, rs.Header, string(body)
}
