package main

import (
	"context"
	"intro/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	log := log.New(os.Stdout, "product-api", log.LstdFlags)

	productHandlers := handlers.NewProducts(log)

	serverMux := http.NewServeMux()
	serverMux.Handle("/", productHandlers)

	server := &http.Server{
		Addr:         ":9000",
		Handler:      serverMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Println("Receive terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)
}
