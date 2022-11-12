package app

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/cleysonph/bookshelf-go/config"
	"github.com/cleysonph/bookshelf-go/internal/infra/app/rest"
)

func Run() {
	var env string
	flag.StringVar(&env, "env", "dev", "Environment to run the application (dev, prod)")
	flag.Parse()

	config.Init(env)

	srv := http.Server{
		Addr:         config.Addr(),
		Handler:      rest.NewRouter(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
