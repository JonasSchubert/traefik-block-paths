package traefik_block_paths_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	BlockPaths "github.com/JonasSchubert/traefik-block-paths"
)

func Test_BlockPaths_ReturnsBlock_IfMatched(t *testing.T) {
	cfg := BlockPaths.CreateConfig()

	cfg.Regex = []string{"^/wp(.*)"}
	cfg.StatusCode = 404

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := BlockPaths.New(ctx, next, cfg, "BlockPaths")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost/wp-login", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("X-Forwarded-For", "2.56.20.0")

	handler.ServeHTTP(recorder, req)

	assertStatusCode(t, recorder.Result(), http.StatusNotFound)
}

func Test_BlockPaths_ReturnsOK_IfNotMatched(t *testing.T) {
	cfg := BlockPaths.CreateConfig()

	cfg.Regex = []string{"^/wp(.*)"}
	cfg.StatusCode = 404

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := BlockPaths.New(ctx, next, cfg, "BlockPaths")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost/index.html", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("X-Forwarded-For", "2.56.20.0")

	handler.ServeHTTP(recorder, req)

	assertStatusCode(t, recorder.Result(), http.StatusOK)
}

func Test_BlockPaths_ReturnsOK_IfMatched_ButLocalIpIsAllowed(t *testing.T) {
	cfg := BlockPaths.CreateConfig()

	cfg.Regex = []string{"^/wp(.*)"}
	cfg.StatusCode = 404

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := BlockPaths.New(ctx, next, cfg, "BlockPaths")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost/wp-login", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("X-Real-IP", "192.168.1.1")

	handler.ServeHTTP(recorder, req)

	assertStatusCode(t, recorder.Result(), http.StatusOK)
}

func Test_BlockPaths_ReturnsBlock_IfMatched_AndLocalIpIsNotAllowed(t *testing.T) {
	cfg := BlockPaths.CreateConfig()

	cfg.AllowLocalRequests = false
	cfg.Regex = []string{"^/wp(.*)"}
	cfg.StatusCode = 404

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := BlockPaths.New(ctx, next, cfg, "BlockPaths")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost/wp-login", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("X-Real-IP", "192.168.1.1")

	handler.ServeHTTP(recorder, req)

	assertStatusCode(t, recorder.Result(), http.StatusNotFound)
}

func assertStatusCode(t *testing.T, req *http.Response, expected int) {
	t.Helper()

	received := req.StatusCode

	if received != expected {
		t.Errorf("invalid status code: %d <> %d", expected, received)
	}
}
