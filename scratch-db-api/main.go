package main

import (
	"log"
	"net/http"
	"scratch-db-api/routes"
)

func main() {

	routes.Startup()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
