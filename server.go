package main

import (
    "net/http"
    "github.com/gorilla/mux"
)

type server struct {
    mux *mux.Router
}

func newServer(mux *mux.Router) *server {
    s := server{mux}
    s.routes() // register handlers
    return &s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    s.mux.ServeHTTP(w, r)
}
