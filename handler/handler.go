package handler

import (
	"net/http"
	"encoding/json"
	"github.com/tiny657/url-shortener/util"
	"github.com/tiny657/url-shortener/storage"
)

type originalParameter struct {
	Url string
}

type shortenParameter struct {
	Short string
}

func HandleShortenUrl(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)

	var parameter originalParameter
	err := decoder.Decode(&parameter)
	if err != nil {
		panic(err)
	}

	shorten := util.MakeShortenUrl(parameter.Url)
	storage.SetUrl(shorten, parameter.Url)

	w.Write(util.CreateShortenResponse(shorten))
}

func HandleOriginalUrl(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)

	var parameter shortenParameter
	err := decoder.Decode(&parameter)
	if err != nil {
		panic(err)
	}

	original := storage.GetUrl(parameter.Short)
	w.Write(util.CreateOriginalResponse(original))
}
