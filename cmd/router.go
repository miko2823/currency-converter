package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/miko2823/currency-converter.git/handler"
	"github.com/miko2823/currency-converter.git/infrastructure/persistence"
	"github.com/miko2823/currency-converter.git/usecase"
)

type Routing struct {
	config Config
}

func (routing *Routing) buildHandler() http.Handler {

	mux := chi.NewRouter()
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorizations", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Use(middleware.Heartbeat("/ping"))
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	converterPersistence := persistence.NewConverterPersistence(routing.config.Env, routing.config.Postgres)
	converterUseCase := usecase.NewConverterUsecase(converterPersistence)
	converterHandler := handler.NewConverterHandler(converterUseCase)

	// mux.Route("/v1", func(r chi.Router) {
	// v1を外にだしたい
	mux.Mount("/converter", converterHandler.RegisterHandlers())

	// Debug
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(mux, walkFunc); err != nil {
		fmt.Printf("Logging err: %s\n", err.Error())
	}

	return mux
}
