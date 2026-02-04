package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pID, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please send a valid product ID", http.StatusBadRequest)
		return
	}

	for _, product := range database.ProductList {
		if product.ID == pID {
			util.SendData(w, product, http.StatusOK)
			return
		}
	}

	util.SendData(w, "Product not found", http.StatusNotFound)
}