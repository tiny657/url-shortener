package util

import "testing"

func TestMakeShortenUrl(t *testing.T) {
	shortenUrl := MakeShortenUrl("http://a.very.long.url")
	if string(shortenUrl) != "http://localhost/ea7c045ff32d2295067299b0f6b214f000114ccb" {
		t.Fail()
	}
}
