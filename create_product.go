package main

import (
	"encoding/json"
	"net/http"
)

func createProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Please sent valid JSON data", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(ProductList) + 1
	ProductList = append(ProductList, newProduct)
	sendData(w, newProduct, 201)
}
