package main

func (s *server) routes() {
    s.mux.HandleFunc("/api/jokes", s.getJokes).Methods("GET")
}
