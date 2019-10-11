package main

import (
    "net/http"
    "net/http/httptest"
    "os"
    "testing"
    "github.com/gorilla/mux"
)

var testServer *server

func TestMain(m *testing.M) {
    testServer = newServer(
        mux.NewRouter(),
    )

    os.Exit(m.Run())
}

func TestEndpointGetJokes(t *testing.T) {
    req, _ := http.NewRequest("GET", "/api/jokes", nil)
    response := executeRequest(req)

    if response.Result().StatusCode != http.StatusOK {
        t.Errorf("Expected an %d status received a %d", http.StatusOK, response.Result().StatusCode)
    }
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
    rr := httptest.NewRecorder()
    testServer.mux.ServeHTTP(rr, req)

    return rr
}
