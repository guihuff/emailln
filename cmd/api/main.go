package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World !!!"))

	})
	fmt.Println("O servidor está rodando na porta 3000")
	http.ListenAndServe(":3000", r)
}
