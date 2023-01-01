package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (p converterHandler) RegisterHandlers() http.Handler {
	r := chi.NewRouter()
	r.Get("/", p.getLatestRates)
	return r
}
