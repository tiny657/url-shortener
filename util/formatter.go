package util

import "encoding/json"

type ShortenResponse struct {
	Short string
}

func CreateShortenResponse(shorten string) []byte {
	response := ShortenResponse{shorten}
	responseString, _ := json.Marshal(response)
	return responseString
}
