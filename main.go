package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request){
		log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})

	http.ListenAndServe(":9000", nil)
}