package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	log *log.Logger
}

func NewHello(log *log.Logger) *Hello {
	return &Hello{log}
}

func (h *Hello) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	h.log.Println("Hello World")

	d, err := ioutil.ReadAll(req.Body)

	if err != nil {
		http.Error(res, "Oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(res, "Hello %s", d)
}
