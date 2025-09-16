package main

import (
	"net/http"

	generated "github.com/bromivipo/marketplace/partner_api/definitions"
	"github.com/bromivipo/marketplace/partner_api/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	handler := generated.NewStrictHandler(handlers.NewServer(), nil)
	generated.HandlerFromMux(handler, router)
	http.ListenAndServe("localhost:8000", router)
}