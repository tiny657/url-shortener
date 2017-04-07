package main

import (
	"net/http"
	"github.com/tiny657/url-shortener/handler"
)

const (
	port = ":8080"
	shortenUrl = "/shorten"
)

func main() {
	http.HandleFunc(shortenUrl, handler.ShortenHandler)
	http.ListenAndServe(port, nil)
}
