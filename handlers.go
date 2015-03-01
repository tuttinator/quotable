package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"html"
	"net/http"
)

type Quote struct {
	Id   int
	Key  string
	Text string
}

func QuoteShowHandler(w http.ResponseWriter, c *RequestContext) {
	key := c.Params["key"]
	fmt.Println(key)
	quote := Quote{}

	err := c.DB.QueryRow("SELECT id, key, text FROM quotes WHERE key=?", key).Scan(&quote.Id, &quote.Key, &quote.Text)
	checkErr(err)

	data, err := json.Marshal(quote)
	checkErr(err)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func ImageServeHandler(w http.ResponseWriter, c *RequestContext) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(c.Request.URL.Path))
}

func QuoteCreateHandler(w http.ResponseWriter, c *RequestContext) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(c.Request.URL.Path))
}
