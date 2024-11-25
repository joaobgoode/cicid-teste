package main

import (
	"fmt"
	"go-app/internals/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/produtos", handlers.GetAllProdutos)
	mux.HandleFunc("GET /api/produtos/{id}", handlers.GetProduto)
	mux.HandleFunc("PUT /api/produtos/{id}", handlers.UpdateProduto)
	mux.HandleFunc("DELETE /api/produtos/{id}", handlers.DeleteProduto)
	mux.HandleFunc("POST /api/new", handlers.NovoProduto)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
