package middleware

import "net/http"

type WrapperWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (w *WrapperWriter) WriteHeader(statusCod int) {
	w.ResponseWriter.WriteHeader(statusCod)
	w.StatusCode = statusCod
}
