package internal

import (
	"item-comparison-api/internal/api"

	"github.com/go-chi/chi/v5"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api/v1/products", func(r chi.Router) {
		r.Get("/", api.ListProducts)
		//r.Get("/{id}", api.GetProduct)
	})

	return r
}
