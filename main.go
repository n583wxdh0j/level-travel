package main

import (
	"log"
	"net/http"
)

const (
	port = ":8080"
)

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", index)

	srv := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(418)
}
