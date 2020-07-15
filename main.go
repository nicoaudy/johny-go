package main

import (
	"intro/handlers"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	log := log.New(os.Stdout, "product-api", log.LstdFlags)

	helloHandlers := handlers.NewHello(log)
	goodbyeHandlers := handlers.NewGoodbye(log)

	serverMux := http.NewServeMux()
	serverMux.Handle("/", helloHandlers)
	serverMux.Handle("/goodbye", goodbyeHandlers)

	server := &http.Server{
		Addr:         ":9000",
		Handler:      serverMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	server.ListenAndServe()
}
