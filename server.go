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

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

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
