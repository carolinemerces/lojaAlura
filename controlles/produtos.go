package controlles

import (
	"html/template"
	"net/http"

	"github.com/carolinemerces/lojaAlura/models"
)

//encapsula os templates e devolve um erro
var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)

}
