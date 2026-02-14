package product

import (
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

	product, err := h.productRepo.Get(pID)
	if err != nil {
		util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if product == nil {
		util.SendError(w, "Product not found", http.StatusNotFound)
		return
	}

	util.SendData(w, product, http.StatusOK)
}