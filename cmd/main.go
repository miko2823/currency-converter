package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/miko2823/currency-converter.git/config"
)

const webPort = "80"

type Config struct {
	Env      config.Environment
	Postgres *sql.DB
}

func main() {

	postgresConn, err := config.ConnectToDB()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	env, err := config.GetEnvironment()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	app := Config{
		Env:      env,
		Postgres: postgresConn,
	}

	var routing = Routing{app}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: routing.buildHandler(),
	}
	log.Printf("Starting service on port %s\n", webPort)

	err = server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
