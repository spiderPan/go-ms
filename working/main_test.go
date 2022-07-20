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

func TestApi(t *testing.T) {
	data := []struct {
		method             string
		endpoint           string
		handler            string
		requestData        string
		expectedStatusCode int
		expectedBody       string
	}{
		{"GET", "/hello", "hello", "pan test", 200, "Hello pan test"},
		{"GET", "/goodbye", "goodbye", "", 200, "Byee"},
	}

	for _, e := range data {
		t.Run(e.endpoint, func(t *testing.T) {
			requestReader := strings.NewReader(e.requestData)

			req := httptest.NewRequest(e.method, e.endpoint, requestReader)
			w := httptest.NewRecorder()
			switch e.handler {
			case "hello":
				hh := handlers.NewHello(logger)
				hh.ServeHTTP(w, req)
			case "goodbye":
				hh := handlers.NewGoodbye(logger)
				hh.ServeHTTP(w, req)

			case "products":
				ph := handlers.NewProducts(logger)
				ph.GetProducts(w, req)
			}

			resp := w.Result()
			body, _ := io.ReadAll(resp.Body)

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("Expect status code to be %v, rather than %v", e.expectedStatusCode, resp.StatusCode)
			}

			if string(body) != e.expectedBody {
				t.Errorf("Expect body to be %v, rather than %v", e.expectedBody, string(body))
			}
		})
	}
}
