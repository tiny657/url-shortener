package server

import (
	"net/http"
	"github.com/tiny657/url-shortener/handler"
)

const (
	port = ":8080"
	originUrl = "/original"
)

func StartServer() {
	http.HandleFunc(originUrl, handler.ShortenHandler);

	http.ListenAndServe(port, nil)
}
