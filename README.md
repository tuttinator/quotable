# Quotable

A webservice, written in Go, to generate Twitter-size images of quotes
from articles as images.

## How does it work?

![Quotable demo](http://i.imgur.com/QBg0DAm.gif)

The frontend JavaScript posts the contents of the current highlighted
selection to the `/create` endpoint.

This generates a unique string as a `key`, which is the URL segment.

`/{key}.png` serves a PNG

`/{key}.json` gives the database row result as JSON

## TODO

There is a chance of key string collisions. There is a uniqueness
validation in `setup.sql`.

We probably shouldn't `panic` on bad JSON, so handle this in a nicer way.

## Contributing

Pull requests welcome
