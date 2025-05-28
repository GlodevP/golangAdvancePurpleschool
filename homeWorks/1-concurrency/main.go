package main

import (
	"fmt"
	"math/rand"
)

func main() {
	sl := make([]int, 10)
	chIn := make(chan int)
	chOut := make(chan int)
	go random10SliceInChanel(chIn)
	go sqrt10ChInChOut(chIn, chOut)
	for i := range sl {
		sl[i] = <-chOut
	}
	close(chOut)
	for i := range sl {
		fmt.Printf("%d ", sl[i])
	}
}

func sqrt10ChInChOut(chIn chan int, chOut chan int) {
	for i := 0; i < 10; i++ {
		j := <-chIn
		chOut <- j * j
	}
	close(chIn)
}

func random10SliceInChanel(ch chan int) {
	sl := make([]int, 10)
	for i := range 10 {
		sl[i] = rand.Intn(101)
	}
	for i := range sl {
		ch <- sl[i]
	}
}
