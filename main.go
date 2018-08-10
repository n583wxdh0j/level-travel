package main

import (
	"level-travel/controllers"
	"log"
	"net/http"
	"time"
)

const (
	port = ":8080"
)

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", controllers.Index)

	srv := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	go func() {
		err := srv.ListenAndServe()
		log.Fatal(err)
	}()

	controllers.UpdateLibrariesData(time.Hour * 24)
}
