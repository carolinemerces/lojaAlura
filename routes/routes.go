package routes

import (
	"net/http"

	"github.com/carolinemerces/lojaAlura/controlles"
)

func CarregaRotas() {
	http.HandleFunc("/", controlles.Index)
	http.HandleFunc("/new", controlles.New)
}
