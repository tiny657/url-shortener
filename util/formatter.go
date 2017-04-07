package util

import "encoding/json"

type ShortenResponse struct {
	Short string
}

type OriginalResponse struct {
	Original string
}

func CreateShortenResponse(shorten string) []byte {
	response := ShortenResponse{shorten}
	responseString, _ := json.Marshal(response)
	return responseString
}

func CreateOriginalResponse(original string) []byte {
	response := OriginalResponse{original}
	responseString, _ := json.Marshal(response)
	return responseString
}
