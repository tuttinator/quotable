package main

func DefineRoutes(s *Server) {
	s.Router.HandleFunc("/{id}.json", QuoteShowHandler).Methods("GET")
	s.Router.HandleFunc("/{id}.jpg", ImageServeHandler).Methods("GET")
	s.Router.HandleFunc("/create", QuoteCreateHandler).Methods("POST")
}
