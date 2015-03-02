package main

import (
	"code.google.com/p/freetype-go/freetype"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
	"strings"
)

type Quote struct {
	Key  string `json:"key"`
	Url  string `json:"url"`
	Text string `json:"text"`
}

func FindQuoteByKey(key string, db *sql.DB) Quote {
	fmt.Println("Database lookup for " + key)
	quote := Quote{}

	err := db.QueryRow("SELECT key, url, text FROM quotes WHERE key=? LIMIT 1", key).Scan(&quote.Key, &quote.Url, &quote.Text)
	checkErr(err)

	return quote
}

func (q *Quote) Image() *image.RGBA {
	textLayer := q.TextToImage()

	pngFile, err := os.Open("./assets/base.png")
	checkErr(err)

	img, err := png.Decode(pngFile)
	checkErr(err)

	c := image.NewRGBA(image.Rect(0, 0, 755, 378))

	// draw bottom layer from file
	draw.Draw(c, c.Bounds(), img, image.Point{0, 0}, draw.Src)
	// draw text layer on top
	draw.Draw(c, c.Bounds(), textLayer, image.Point{0, 0}, draw.Over)

	return c
}

func (q *Quote) TextToImage() *image.RGBA {
	spacing := 1.5
	var fontsize float64 = 8

	// read font
	fontBytes, err := ioutil.ReadFile(os.Getenv("FONT_FILE"))
	checkErr(err)
	font, err := freetype.ParseFont(fontBytes)
	checkErr(err)

	// Initialize the context.
	fg, bg := image.White, image.Transparent

	// 755px by 378px is the size Vox uses
	rgba := image.NewRGBA(image.Rect(0, 0, 755, 378))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(300)
	c.SetFont(font)
	c.SetFontSize(fontsize)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)

	// Draw the text
	pt := freetype.Pt(50, 50+int(c.PointToFix32(fontsize)>>8))
	lines := strings.Split(q.Text, "\n")
	for _, s := range lines {
		_, err = c.DrawString(s, pt)
		checkErr(err)
		pt.Y += c.PointToFix32(fontsize * spacing)
	}

	return rgba
}
