package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	t := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {

			getHTTPCode("https://ya.ru")
			wg.Done()
		}()
	}
	wg.Wait()
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
