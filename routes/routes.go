package routes

import (
	"net/http"

	"github.com/carolinemerces/lojaAlura/controlles"
)

func CarregaRotas() {
	http.HandleFunc("/", controlles.Index)
	http.HandleFunc("/new", controlles.New)
	http.HandleFunc("/insert", controlles.Insert)
	http.HandleFunc("/delete", controlles.Delete)
	http.HandleFunc("/edit", controlles.Edit)
}

