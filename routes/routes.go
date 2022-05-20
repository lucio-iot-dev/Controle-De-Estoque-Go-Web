package routes

import (
	"Controle-De-Estoque-Go-Web/controllers"
	"net/http"
) 

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
