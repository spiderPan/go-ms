package main_test

import (
	"go-ms/handlers"
	"io"
	"log"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var logger *log.Logger

func TestMain(m *testing.M) {
	logger = log.New(os.Stdout, "product-api", log.LstdFlags)
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestHello(t *testing.T) {
	hh := handlers.NewHello(logger)
	name := "pan test"

	req := httptest.NewRequest("GET", "/hello", strings.NewReader(name))
	w := httptest.NewRecorder()
	hh.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	expectResponseBody := "Hello " + name

	if resp.StatusCode != 200 {
		t.Errorf("Status Code error")
	}

	if string(body) != expectResponseBody {
		t.Errorf("Expect body to be %v, rather than %v", expectResponseBody, string(body))
	}
}

func TestGoodbye(t *testing.T) {
	hh := handlers.NewGoodbye(logger)

	req := httptest.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()
	hh.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	expectResponseBody := "Byee"

	if resp.StatusCode != 200 {
		t.Errorf("Status Code error")
	}

	if string(body) != expectResponseBody {
		t.Errorf("Expect body to be %v, rather than %v", expectResponseBody, string(body))
	}
}
