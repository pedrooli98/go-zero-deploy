package main

import (
	"fmt"
	"net/http"
)

func main() {

	handler := Handler{}
	http.ListenAndServe(":5000", handler)
	fmt.Println("Servidor rodando na porta 5000")
}

type Handler struct{}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
