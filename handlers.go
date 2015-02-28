package main

import (
	"database/sql"
	"fmt"
	"html"
	"net/http"

	"github.com/gorilla/mux"
)

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
