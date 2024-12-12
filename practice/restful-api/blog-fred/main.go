package main

import (
	"log"
	"net/http"

	"project01/practice/restful-api/blog-fred/route"
)

func main() {
	r := route.NewRouter() // initialize the router
	log.Fatal(http.ListenAndServe(":8080", r))
}
