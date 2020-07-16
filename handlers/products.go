package handlers

import (
	"encoding/json"
	"intro/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l*log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	productList := data.GetProducts()
	data, err := json.Marshal(productList)

	if err != nil {
		http.Error(res, "Unable to marshal json", http.StatusInternalServerError)
	}

	res.Write(data)
}
