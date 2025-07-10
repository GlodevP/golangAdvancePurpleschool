package middleware

import (
	"log"
	"net/http"
	"strings"
)

func IsAuthed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authString := r.Header.Get("Authorization")
		tocken := strings.TrimPrefix(authString,"Bearer ")
		log.Println(tocken)
		next.ServeHTTP(w, r)
	})
}