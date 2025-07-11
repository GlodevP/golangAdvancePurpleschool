package middleware

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := &WrapperWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapper, r)
		end := time.Since(start)
		log.WithFields(log.Fields{
			"status":wrapper.StatusCode,
			"method":r.Method,
			"url":r.URL,
			"time":end,
		}).Info()
	})
}
