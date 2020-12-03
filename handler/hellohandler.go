package handler

import "net/http"

// HelloHandler writes a Hello World! response
type HelloHandler struct {
}

// NewHelloHandler instantiates a HelloHandler
func NewHelloHandler() HelloHandler {
	return HelloHandler{}
}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte("Hello World! This is an update"))
	w.WriteHeader(http.StatusOK)
}
