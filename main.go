package main

import (
	"models/routes"
	"net/http"
)

func main() {
	routes.Rotas()
	http.ListenAndServe("127.0.0.1:4000", routes.Mux)
}
