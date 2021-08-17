package main

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/olusolaa/go-todo/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r :=mux.NewRouter()
	routes.TodoRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
