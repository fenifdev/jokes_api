package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
    fmt.Printf("hello, world\n")
    mux := mux.NewRouter()
    server := newServer(mux)
    http.ListenAndServe(":8080", server)
}
