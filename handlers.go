package main

import (
	"bufio"
	"encoding/json"
	"image/png"
	"net/http"
	"time"
)

func QuoteShowHandler(w http.ResponseWriter, c *RequestContext) {
	quote := FindQuoteByKey(c.Params["key"], c.DB)

	data, err := json.Marshal(quote)
	checkErr(err)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func ImageServeHandler(w http.ResponseWriter, c *RequestContext) {
	w.Header().Set("Content-Type", "image/png")
	// figure out the Content-Length of this new combined image somehow
	// w.Header().Set("Content-Length", fmt.Sprint(pngImage.ContentLength))

	quote := FindQuoteByKey(c.Params["key"], c.DB)

	img := quote.Image()
	b := bufio.NewWriter(w)
	png.Encode(b, img)

}

func QuoteCreateHandler(w http.ResponseWriter, c *RequestContext) {
	key := KeyGenerator(10)
	quote := Quote{}

	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&quote)
	checkErr(err)

	quote.Key = key

	stmt, err := c.DB.Prepare("INSERT INTO quotes (key, url, text, created_at) VALUES ($1, $2, $3, $4);")
	checkErr(err)
	_, err = stmt.Exec(quote.Key, quote.Url, quote.Text, time.Now())
	checkErr(err)

	data, err := json.Marshal(quote)
	checkErr(err)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
