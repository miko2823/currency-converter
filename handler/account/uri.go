package account

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h accountHandler) RegisterHandlers() http.Handler {
	r := chi.NewRouter()
	r.Post("/login", h.Login)
	return r
}
