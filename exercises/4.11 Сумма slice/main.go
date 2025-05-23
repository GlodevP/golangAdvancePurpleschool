package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	numGoroutines := 3
	ch := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			last := (i * 4) + 4
			last = min(last, len(arr))
			summSlice(arr[i*4:last], ch)
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	var summ int
	for i := range ch {
		summ += i
	}
	fmt.Printf("Сумма слайса равна: %d\n", summ)
}
func summSlice(arr []int, ch chan int) {
	var summ int
	for _, i := range arr {
		summ += i
	}
	ch <- summ
}
