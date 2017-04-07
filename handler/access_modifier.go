package handler

import (
	"net/http"
)

type handler func(w http.ResponseWriter, r *http.Request)

func GetOnly(handler handler) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handler(w, r)
			return
		}
		http.Error(w, "Only GET method is allowed.", http.StatusMethodNotAllowed)
	}
}

func PostOnly(handler handler) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			handler(w, r)
			return
		}
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
	}
}
