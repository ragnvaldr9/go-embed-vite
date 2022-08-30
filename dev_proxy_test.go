package vite

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

const PROXY_TARGET_HOST = "http://localhost"
const PROXY_TARGET_PORT = "8091"

var testServer *http.Server

const TEST_RES_DATA = "Hello, world!"

func runTestServer() {
	m := http.NewServeMux()

	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(TEST_RES_DATA))
	})

	testServer = &http.Server{
		Addr:    ":" + PROXY_TARGET_PORT,
		Handler: m,
	}

	log.Fatal(testServer.ListenAndServe())
}

func TestDevProx(t *testing.T) {
	go runTestServer()

	proxy := &DevProxy{url: fmt.Sprintf("%v:%v", PROXY_TARGET_HOST, PROXY_TARGET_PORT)}

	handler := http.StripPrefix("/", proxy)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	res := w.Result()

	defer res.Body.Close()
	defer testServer.Close()

	data, _ := ioutil.ReadAll(res.Body)

	if string(data) != TEST_RES_DATA {
		t.Errorf("expected %v got %v", TEST_RES_DATA, string(data))
	}
}
