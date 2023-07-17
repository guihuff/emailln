package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type product struct {
	Id   int
	Name string
}

func main() {
	r := chi.NewRouter()
	r.Use(myMiddleware)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Query().Get("name")
		if param != "" {
			w.Write([]byte(param))
		} else {
			w.Write([]byte("sem parametro"))
		}
	})

	r.Get("/{productName}", func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "productName")
		w.Write([]byte(param))
	})

	r.Get("/json", func(w http.ResponseWriter, r *http.Request) {
		obj := map[string]string{"message": "sucess"}
		render.JSON(w, r, obj)
	})

	r.Post("/product", func(w http.ResponseWriter, r *http.Request) {
		var product product
		render.DecodeJSON(r.Body, &product)
		product.Id = 5
		render.JSON(w, r, product)
	})

	fmt.Println("Porta: 3000")
	http.ListenAndServe(":3000", r)

}

func myMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("before")
		next.ServeHTTP(w, r)
		println("after")
	})
}

