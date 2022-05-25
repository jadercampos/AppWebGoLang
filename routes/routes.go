package routes

import (
	"net/http"

	"appwebgolang.com/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
