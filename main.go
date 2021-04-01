package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

//encapsula os templates e devolve um erro
var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{"Camiseta", "Azul", 49.99, 6},
		{"Tênis", "Preto", 79.98, 7},
		{"Fone", "Sem fio", 129.99, 9},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}
