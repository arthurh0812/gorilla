package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/arthurh0812/gorilla/pkg/handlers"
)

func main() {
	r := mux.NewRouter()
	logger := log.Default()

	root := r.Path("/")
	root.Name("rootRoute")
	root.Handler(handlers.NewRoot(logger))

	api := r.Path("/api")
	api.Name("apiRoute")
	api.Handler(handlers.NewAPI(logger))

	login := r.Path("/login")
	login.Name("loginRoute")
	login.Handler(handlers.NewLogin(logger))

	r.NotFoundHandler = handlers.NewError(logger, http.StatusNotFound)

	err := http.ListenAndServe(Env["HOST"] + ":" + Env["PORT"], r)
	if err != nil {
		log.Fatal(err)
	}
}
