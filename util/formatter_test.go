package util

import (
	"testing"
)

func TestCreateShortenResponse(t *testing.T) {
	shorten := CreateShortenResponse("http://localhost/1x2x3x")
	if string(shorten) != "{\"Short\":\"http://localhost/1x2x3x\"}" {
		t.Fail()
	}
}
