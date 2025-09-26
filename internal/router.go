package internal

import (
	"log"
	"net/http"
	"time"

	"item-comparison-api/internal/api"

	"github.com/go-chi/chi/v5"
)

// loggingMiddleware logs each incoming HTTP request
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s %s in %v", r.Method, r.RequestURI, time.Since(start))
	})
}

func SetupRouter(handler *api.ProductHandler) *chi.Mux {
	// Create a new router
	r := chi.NewRouter()

	// Apply logging middleware
	r.Use(loggingMiddleware)

	// Define routes
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
