package converter

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (p converterHandler) RegisterHandlers() http.Handler {
	r := chi.NewRouter()
	r.Get("/", p.getLatestRates)
	r.Get("/get_all_symbols", p.getAllSymbols)
	return r
}
