package main

import "net/http"

func DefineRoutes(s *Server) {
	s.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/index.html")
	}).Methods("GET")
	s.DefineRoute("/{key}.json", QuoteShowHandler).Methods("GET")
	s.DefineRoute("/{key}.png", ImageServeHandler).Methods("GET")
	s.DefineRoute("/create", QuoteCreateHandler).Methods("POST")
	s.Router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./public/"))))
}
