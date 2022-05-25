package main

import (
	"net/http"

	"appwebgolang.com/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8100", nil)
}
