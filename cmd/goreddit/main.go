package main

import (
	"log"
	"net/http"

	"github.com/raynine/goreddit/postgres"
	"github.com/raynine/goreddit/web"
)

func main() {
	store, err := postgres.NewStore("postgresql://postgres:admin@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	handler := web.NewHandler(store)
	http.ListenAndServe(":3000", handler)
}
