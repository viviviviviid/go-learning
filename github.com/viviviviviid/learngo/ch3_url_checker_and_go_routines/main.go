package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
		"https://error.erorrorroro.com",
	}
	// 두가지 값이 나오는데, 첫번째는 인덱스, 두번째가 내용
	for _, url := range urls {
		check := hitURL(url)
		if check != nil {
			fmt.Println(url, ": ", check)
		}
	}
}

func hitURL(url string) error {
	// hit : 특정 웹사이트의 파일 1개에 접속하는 것.
	fmt.Println("Checking: ", url)
	// golang standard package // https://pkg.go.dev/net/http@go1.20.5
	// resp, err := http.Get("http://example.com/")
	// resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
	// resp, err := http.PostForm("http://example.com/form", url.Values{"key": {"Value"}, "id": {"123"}})
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
