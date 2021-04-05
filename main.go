package main

import (
	"net/http"
	"text/template"

	"github.com/carolinemerces/lojaAlura/models"
)

//encapsula os templates e devolve um erro
var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
	
}
