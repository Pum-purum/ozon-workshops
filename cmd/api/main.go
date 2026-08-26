package main

import (
	"log"
	"net/http"

	product "shop/internal/product"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /products", product.ListHandler)
	mux.HandleFunc("GET /products/{id}", product.GetHandler)

	log.Println("listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
