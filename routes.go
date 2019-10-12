package main

func (s *server) routes() {
    s.mux.HandleFunc("/api/jokes", s.getJokes).Methods("GET")
    s.mux.HandleFunc("/api/jokes/{id:[0-9]+}", s.getJokeByID).Methods("GET")
    s.mux.HandleFunc("/api/jokes", s.postJokes).Methods("POST")
    s.mux.HandleFunc("/api/jokes/{id:[0-9]+}", s.deleteJokeByID).Methods("DELETE")
}
