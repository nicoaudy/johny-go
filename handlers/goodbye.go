package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	log *log.Logger
}

func NewGoodbye(log *log.Logger) *Goodbye {
	return &Goodbye{log}
}

func (g *Goodbye) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	g.log.Println("logger Goodbye")

	res.Write([]byte("Bye Bruh"))
}
