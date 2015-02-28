package main

import (
	"database/sql"
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

	DefineRoutes(server)

	return server
}

func (s *Server) Close() {
	s.DB.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
