package main

import (
	"fmt"
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
	fmt.Println(s.ListenAndServe())
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	randInt := rand.Int31n(6) + 1
	randIntStr := strconv.Itoa(int(randInt))
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(randIntStr))
}
