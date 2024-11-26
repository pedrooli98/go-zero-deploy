package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http.Handle("/", Handler{})
	// http.Handle("/hello", HelloHandler{})
	// http.Handle("/world", WorldHandler{})
	handler := GenericHandler{}

	http.ListenAndServe(":5000", handler)
	fmt.Println("Servidor rodando na porta 5000")
}

type Handler struct{}
type HelloHandler struct{}
type WorldHandler struct{}
type GenericHandler struct{}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func (h WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("World"))
}

func (h GenericHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/hello":
		w.Write([]byte("Hello"))
	case "/world":
		w.Write([]byte("World"))
	default:
		w.Write([]byte("404 Not Found"))
	}
}
