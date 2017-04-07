package main

import (
	"net/http"
	"github.com/tiny657/url-shortener/util"
)

const (
	port = ":8080"
	shortenUrl = "/shorten"
)

func main() {
	http.HandleFunc(shortenUrl, util.ShortenHandler)
	http.ListenAndServe(port, nil)
}
