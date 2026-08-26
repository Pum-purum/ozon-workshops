package main

import (
	"log"
	"net/http"

	"shop/internal/cart"
	"shop/internal/cart/productclient"
)

func main() {
	productClient := productclient.New("http://localhost:8080")
	service := cart.NewService(productClient)
	handler := cart.NewHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /user/{user_id}/cart/{sku_id}", handler.Add)
	mux.HandleFunc("DELETE /user/{user_id}/cart/{sku_id}", handler.Remove)
	mux.HandleFunc("DELETE /user/{user_id}/cart", handler.Clear)
	mux.HandleFunc("GET /user/{user_id}/cart", handler.Get)

	log.Println("cart service listening on :8082")
	if err := http.ListenAndServe(":8082", mux); err != nil {
		log.Fatal(err)
	}
}
