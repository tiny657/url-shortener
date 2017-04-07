package util

import (
	"net/http"
	"io/ioutil"
)

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var original, _ = ioutil.ReadAll(r.Body)

	var shorten = MakeShortenUrl(string(original))
	w.Write([]byte(shorten))
}
