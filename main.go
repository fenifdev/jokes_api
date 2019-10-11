package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
  	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Joke{})

    fmt.Printf("hello, world\n")

    mux := mux.NewRouter()
    server := newServer(mux, db)
    http.ListenAndServe(":8080", server)
}
