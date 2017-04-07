package server

import (
	"net/http"
	"github.com/tiny657/url-shortener/handler"
)

const (
	port = ":8080"
	originUrl = "/original"
	shortenUrl = "/shorten"
)

func StartServer() {
	http.HandleFunc(shortenUrl, handler.HandleShortenUrl);
	http.HandleFunc(originUrl, handler.HandleOriginalUrl);

	http.ListenAndServe(port, nil)
}
