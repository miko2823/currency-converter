package converter

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (p converterHandler) RegisterHandlers() http.Handler {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://d1m0p5gfo03e5l.cloudfront.net", "https://agiicorp.net"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Requested-With"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	r := chi.NewRouter()
	r.Get("/", p.getLatestRates)
	r.Get("/get_all_symbols", p.getAllSymbols)
	return r
}
