package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	account_handler "github.com/miko2823/currency-converter.git/handler/account"
	converter_handler "github.com/miko2823/currency-converter.git/handler/converter"
	account_infra "github.com/miko2823/currency-converter.git/infrastructure/account"
	converter_infra "github.com/miko2823/currency-converter.git/infrastructure/converter"
	account_usecase "github.com/miko2823/currency-converter.git/usecase/account"
	converter_usecase "github.com/miko2823/currency-converter.git/usecase/converter"
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

	accountPersistence := account_infra.NewAccountPersistence(routing.config.Env, routing.config.Postgres)
	accountUseCase := account_usecase.NewAccountUsecase(routing.config.Env, accountPersistence)
	accountHandler := account_handler.NewAccountHandler(accountUseCase)

	converterPersistence := converter_infra.NewConverterPersistence(routing.config.Env, routing.config.Postgres)
	converterUseCase := converter_usecase.NewConverterUsecase(converterPersistence)
	converterHandler := converter_handler.NewConverterHandler(converterUseCase)

	mux.Mount("/", accountHandler.RegisterHandlers())
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
