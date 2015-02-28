package main

func DefineRoutes(s *Server) {
	s.DefineRoute("/{key}.json", QuoteShowHandler).Methods("GET")
	s.DefineRoute("/{key}.png", ImageServeHandler).Methods("GET")
	s.DefineRoute("/create", QuoteCreateHandler).Methods("POST")
}
