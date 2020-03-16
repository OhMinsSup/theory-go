package main

import (
	"errors"
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

var errorRequestFailed = errors.New("Request failed")

func main() {
	c := make(chan result)
	urls := []string{
		"https://github.com/",
		"https://www.instagram.com/?hl=ko",
		"https://www.google.co.kr/",
		"https://www.typescriptlang.org/",
		"https://www.docker.com/",
		"https://ko.reactjs.org/",
		"https://flutter.dev/",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<- c)
	}
}

func hitURL(url string, c chan<- result) {
	res, err := http.Get(url)
	status := "OK"
	if err != nil || res.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- result{
		url:    url,
		status: status,
	}
}
