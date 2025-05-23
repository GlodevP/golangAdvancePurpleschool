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
	code := make(chan int)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			getHTTPCode("https://ya.ru", code)
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(code)
	}()
	for statusCode := range code {
		fmt.Printf("Код ответа: %d\n", statusCode)
	}

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
