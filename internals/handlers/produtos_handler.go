package handlers

import (
	"fmt"
	"net/http"
)

func GetAllProdutos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Listar todos os produtos")
}

func GetProduto(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "Obter produto %s\n", id)
}

func NovoProduto(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Adicionar novo produto")
}

func UpdateProduto(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "Atualizar produto %s\n", id)
}

func DeleteProduto(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "Deletar produto %s\n", id)
}
