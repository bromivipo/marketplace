package main

import (
	"log"
	"net/http"

	generated "github.com/bromivipo/marketplace/partner_api/definitions"
	"github.com/bromivipo/marketplace/partner_api/handlers"
	"github.com/bromivipo/marketplace/partner_api/pgrepo"
	"github.com/go-chi/chi/v5"
)

func main() {
	log.Println("Partner Service started")
	router := chi.NewRouter()
	handler := generated.NewStrictHandler(handlers.NewServer(), nil)
	generated.HandlerFromMux(handler, router)
	http.ListenAndServe(pgrepo.GetEnvOrDefault("PARTNER_API_HTTP_ADDRESS", "localhost:8000"), router)
}