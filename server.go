package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/jinzhu/gorm"
)

type server struct {
    mux *mux.Router
    db *gorm.DB
}

func newServer(mux *mux.Router, db *gorm.DB) *server {
    s := server{mux, db}
    s.routes() // register handlers
    return &s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    s.mux.ServeHTTP(w, r)
}
