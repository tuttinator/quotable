package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"sync"

	"github.com/gorilla/mux"
)

type Server struct {
	DB        *sql.DB
	Router    *mux.Router
	waitGroup *sync.WaitGroup
}

type HandlerWithContext func(
	http.ResponseWriter,
	*RequestContext)

type RequestContext struct {
	DB      *sql.DB
	Params  map[string]string
	Request *http.Request
}

func NewServer() *Server {

	server := &Server{
		SetupDB(),
		mux.NewRouter(),
		&sync.WaitGroup{},
	}

	DefineRoutes(server)

	return server
}

func (s *Server) Close() {
	s.DB.Close()
}

func (s *Server) DefineRoute(pattern string, handler HandlerWithContext) *mux.Route {
	internalHandler := func(w http.ResponseWriter, r *http.Request) {
		s.waitGroup.Add(1)
		defer s.waitGroup.Done()

		defer func() {
			if err := recover(); err != nil {
				log.Printf("PANIC: %s\n%s", err, debug.Stack())
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		context := newRequestContext(s.DB, r)

		handler(w, context)
	}

	return s.Router.HandleFunc(pattern, internalHandler)
}

func newRequestContext(db *sql.DB, r *http.Request) *RequestContext {
	params := mux.Vars(r)
	return &RequestContext{db, params, r}
}

func SetupDB() *sql.DB {

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	checkErr(err)

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS "quotes" (
		"id" SERIAL PRIMARY KEY,
		"key" VARCHAR(64) NOT NULL UNIQUE,
		"url" TEXT NOT NULL,
		"text" TEXT  NOT NULL,
		"created_at" TIMESTAMP NOT NULL
	);
	`)

	checkErr(err)

	return db

}
