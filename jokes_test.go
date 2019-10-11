package main

import (
    "net/http"
    "net/http/httptest"
    "os"
    "testing"
    "github.com/gorilla/mux"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var testServer *server

func TestMain(m *testing.M) {
    db, err := gorm.Open("sqlite3", "./test.db")
    if err != nil {
        panic("failed to connect database")
    }
    defer db.Close()

    // Migrate the schema
    db.AutoMigrate(&Joke{})

    testServer = newServer(mux.NewRouter(), db)

    os.Exit(m.Run())
}

func TestEndpointGetJokes(t *testing.T) {
    req, _ := http.NewRequest("GET", "/api/jokes", nil)
    response := executeRequest(req)

    if response.Result().StatusCode != http.StatusOK {
        t.Errorf("Expected an %d status received a %d", http.StatusOK, response.Result().StatusCode)
    }
}

func TestEndpointGetJokesEmpty(t *testing.T) {
    req, _ := http.NewRequest("GET", "/api/jokes", nil)
    response := executeRequest(req)

    //Expect a 200 status.
    if response.Result().StatusCode != http.StatusOK {
        t.Errorf("Expected an %d status received a %d", http.StatusOK, response.Result().StatusCode)
    }

    //Expect a json type response.
    if content_type := response.Result().Header.Get("Content-Type"); content_type != "application/json" {
        t.Errorf("Expected a content type application/json. Got %s", content_type)
    }
    //Expect an empty array.

    if body := response.Body.String(); body != "[]" {
        t.Errorf("Expected an empty array. Got %s", body)
    }
}

func TestEndpointGetJokesWithResults(t *testing.T) {
    req, _ := http.NewRequest("GET", "/api/jokes", nil)
    response := executeRequest(req)

    if response.Result().StatusCode != http.StatusOK {
        t.Errorf("Expected an %d status received a %d", http.StatusOK, response.Result().StatusCode)
    }
}

func TestEndpointGetJokesById(t *testing.T) {
    req, _ := http.NewRequest("GET", "/api/jokes/1", nil)
    response := executeRequest(req)

    if response.Result().StatusCode != http.StatusOK {
        t.Errorf("Expected an %d status received a %d", http.StatusOK, response.Result().StatusCode)
    }
}

func TestEndpointPostJokes(t *testing.T) {
    req, _ := http.NewRequest("POST", "/api/jokes", nil)
    response := executeRequest(req)

    if response.Result().StatusCode != http.StatusOK {
        t.Errorf("Expected an %d status received a %d", http.StatusOK, response.Result().StatusCode)
    }
}

func TestEndpointDeleteJokesById(t *testing.T) {
    req, _ := http.NewRequest("DELETE", "/api/jokes/1", nil)
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
