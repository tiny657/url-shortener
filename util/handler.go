package util

import (
	"net/http"
	"encoding/json"
)

type originalParameter struct {
	Url string
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)

	var parameter originalParameter
	err := decoder.Decode(&parameter)
	if err != nil {
		panic(err)
	}

	shorten := MakeShortenUrl(parameter.Url)

	w.Write(CreateShortenResponse(shorten))
}


