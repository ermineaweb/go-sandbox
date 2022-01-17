package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sandbox/data"
)

type Product struct {
	logger *log.Logger
}

func NewProduct() Product {
	return Product{logger: log.New(os.Stdout, "product-api-", log.LstdFlags)}
}

func (p Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("request on product handler")
	p.getProducts(rw, r)
}

func (p Product) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("get products")
	json := json.NewEncoder(rw)
	json.Encode(data.ProductList)
}
