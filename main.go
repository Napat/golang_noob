package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Please set PORT env")
	}
	r := mux.NewRouter()
	r.HandleFunc("/", HomesHandler)
	r.HandleFunc("/products/{product_id}", ProductsHandler)

	fmt.Printf("Starting at %s...\n", port)
	http.ListenAndServe(":"+port, r)
}

// HomesHandler Please try http://127.0.0.1:8080/
func HomesHandler(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("content-type", "application/json")
	fmt.Fprintf(w, `{"Hello": "World"}`)
}

// ProductsHandler Please try http://127.0.0.1:8080/products/123
func ProductsHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	productID := vars["product_id"]

	w.Header().Set("content-type", "application/json")
	fmt.Fprintf(w, `{"product_id": "%s"}`, productID)
}
