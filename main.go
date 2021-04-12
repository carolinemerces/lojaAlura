package main

import (
	"net/http"

	"github.com/carolinemerces/lojaAlura/routes"
)


func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
