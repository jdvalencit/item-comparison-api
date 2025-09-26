package internal

import (
	"item-comparison-api/internal/api"

	"github.com/go-chi/chi/v5"
)

func SetupRouter(handler *api.ProductHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api/v1/products", func(r chi.Router) {
		r.Get("/", handler.LoadProducts)
		r.Post("/", handler.SaveProducts)
		r.Put("/", handler.UpdateProducts)
		r.Delete("/{id}", handler.DeleteProduct)
		r.Get("/{id}", handler.GetProduct)
		r.Post("/compare", handler.CompareProducts)
	})

	return r
}
