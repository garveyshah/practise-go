package main

import (
	"log"
	"net/http"

	"api1/route"
)
func main() {
	r := route.NewRouter() // initialize the router
	log.Fatal(http.ListenAndServe(":8080", r))

}