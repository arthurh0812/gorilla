package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/arthurh0812/gorilla/pkg/handlers"
)

func main() {
	r := mux.NewRouter()

	r.Handle("/api", handlers.API{})

	err := http.ListenAndServe(Env["HOST"] + ":" + Env["PORT"], r)
	if err != nil {
		log.Fatal(err)
	}
}
