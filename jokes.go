package main

import(
	"net/http"
	"github.com/jinzhu/gorm"
	"encoding/json"
)

type Joke struct {
	gorm.Model
	Text string
}

func (s *server) getJokes(w http.ResponseWriter, r *http.Request) {
	var jokes[] Joke
	s.db.Find(&jokes)
	payload := jokes
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (s *server) getJokeByID(w http.ResponseWriter, r *http.Request) {

}

func (s *server) postJokes(w http.ResponseWriter, r *http.Request) {

}

func (s *server) deleteJokeByID(w http.ResponseWriter, r *http.Request) {

}
