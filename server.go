package main

import (
	"gocnpj/handlers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8085"
	}

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/cnpj/{cnpj}", handlers.CnpjHandler)

	http.ListenAndServe(":"+port, r)

}
