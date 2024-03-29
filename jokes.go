package main

import(
	"net/http"
	"github.com/jinzhu/gorm"
	"encoding/json"
	"github.com/gorilla/mux"
	//"log"
	//"io/ioutil"
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
	vars := mux.Vars(r)
	var joke Joke

	if err := s.db.First(&joke, vars["id"]).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			response, _ := json.Marshal(err)
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
		}
	} else {
		payload := joke
		response, _ := json.Marshal(payload)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}

}

func (s *server) postJokes(w http.ResponseWriter, r *http.Request) {
	var joke Joke
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		response, _ := json.Marshal(map[string]string{"error": "text is required"})
		w.Write(response)
		return
	}

	decoder := json.NewDecoder(r.Body)

    err := decoder.Decode(&joke)
    if err != nil {
        response, _ := json.Marshal(map[string]string{"error": "text is required"})
		w.Write(response)
	} else {
		s.db.Create(&joke)
		response, _ := json.Marshal(joke)
		w.Write(response)
	}
}

func (s *server) deleteJokeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var joke Joke

	if err := s.db.First(&joke, vars["id"]).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			response, _ := json.Marshal(err)
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
		}
	} else {
		s.db.Delete(&joke)
		payload := joke
		response, _ := json.Marshal(payload)
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}
