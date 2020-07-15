package main

import (
	"intro/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	log := log.New(os.Stdout, "product-api", log.LstdFlags)

	helloHandlers := handlers.NewHello(log)
	goodbyeHandlers := handlers.NewGoodbye(log)

	serverMux := http.NewServeMux()
	serverMux.Handle("/", helloHandlers)
	serverMux.Handle("/goodbye", goodbyeHandlers)

	http.ListenAndServe(":9000", serverMux)
}
