package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	t := time.Now()
	code := make(chan int)
	go getHTTPCode("https://ya.ru", code)
	fmt.Println(<-code)
	fmt.Println(time.Since(t))
}

func getHTTPCode(url string, codeCh chan int) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error http get from url: %s, error: %s", url, err)
		return
	}
	codeCh <- resp.StatusCode
}
