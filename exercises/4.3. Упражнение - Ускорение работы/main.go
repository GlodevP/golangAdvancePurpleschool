package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	t := time.Now()
	for i := 0; i < 10; i++ {
		go getHTTPCode("https://ya.ru")
	}
	time.Sleep(time.Second)
	fmt.Println(time.Since(t))
}

func getHTTPCode(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error http get from url: %s, error: %s", url, err)
		return
	}
	fmt.Printf("Code: %d \n", resp.StatusCode)

}
