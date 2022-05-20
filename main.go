package main

import (
	"Controle-De-Estoque-Go-Web/routes"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
