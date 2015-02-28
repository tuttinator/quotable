package main

import (
	"database/sql"
	"fmt"
	"html"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Server struct {
	DB        *sql.DB
	Router    *mux.Router
	waitGroup *sync.WaitGroup
}

func NewServer() *Server {
	db, err := sql.Open("sqlite3", "./database.db")
	checkErr(err)

	server := &Server{
		db,
		mux.NewRouter(),
		&sync.WaitGroup{},
	}

	server.Router.HandlerFunc("/{key}.json", QuoteShowHandler).Methods("GET")
	server.Router.HandlerFunc("/{key}.png", ImageServeHandler).Methods("GET")
	server.Router.HandlerFunc("/create", QuoteCreateHandler).Methods("POST")

	DefineRoutes(server)

	return server
}

func (s *Server) Close() {
	s.DB.Close()
}

type Quote struct {
	Key  string
	Text string
}

func QuoteShowHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	params := mux.Vars(r)
	key := params["key"]

	fmt.Fprintf(w, "Key is %q", key)
}

func ImageServeHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func QuoteCreateHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
