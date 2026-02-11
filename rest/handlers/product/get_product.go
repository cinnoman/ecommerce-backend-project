package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pID, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please send a valid product ID", http.StatusBadRequest)
		return
	}

	product := database.GetProduct(pID)
	if product == nil {
		util.SendError(w, "Product not found", http.StatusNotFound)
	}

	util.SendData(w, product, http.StatusOK)
}