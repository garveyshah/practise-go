package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(helloWorld))
	log.Fatal(http.ListenAndServe(":8080", mux))
	fmt.Println("Server RUnning on port : https//localhost:8080")
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "test/plain")

	fmt.Fprintln(w, "Hello, world")
}
