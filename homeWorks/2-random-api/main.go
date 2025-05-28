package main

import (
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", indexHandler)
	s := http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	s.ListenAndServe()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	randInt := rand.Int31n(6) + 1
	randIntStr := strconv.Itoa(int(randInt))
	w.Write([]byte(randIntStr))
}
