package main

import (
	"net/http"

	"github.com/OsvaldoLlNava/dockercises1/structured/app/pkg/DockerciseEndpoints"
	"github.com/go-chi/chi/v5"
)

func handleRequests() {
	r := chi.NewRouter()
	r.Get("/", DockerciseEndpoints.HomePage)
	r.Get("/people", DockerciseEndpoints.AllResults)
	r.Get("/people/{userId}", DockerciseEndpoints.SpecificResult)
	http.ListenAndServe(":7777", r)
}

func main() {
	handleRequests()
}
