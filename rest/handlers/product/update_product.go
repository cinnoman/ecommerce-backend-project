package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	pID, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please send a valid product ID", http.StatusBadRequest)
		return
	}

	var updatedProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Please sent valid JSON data", http.StatusBadRequest)
		return
	}

	updatedProduct.ID = pID

	database.UpdateProduct(updatedProduct)
	util.SendData(w, "Updated successfully", http.StatusCreated)
}
