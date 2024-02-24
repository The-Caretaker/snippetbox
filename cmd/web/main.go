package main

import (
	"log"
	"net/http"
)

func main() {
	//Server
	mux := http.NewServeMux()

	// Routes (Handler Functions)
	mux.HandleFunc("/", home)
	mux.HandleFunc("/calculate", calculate)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :4000")

	// Listening and serving
	err := http.ListenAndServe("localhost:4000", mux)
	log.Fatal(err)
}
