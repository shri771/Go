package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "1030"

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	// Handlers
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./assets"))))
	mux.HandleFunc("GET /echo", handlerEchoMessage)
	mux.HandleFunc("POST /data", handlerAddJson)

	log.Printf("Listing on port: %v \n", port)
	log.Fatal(srv.ListenAndServe())
}
