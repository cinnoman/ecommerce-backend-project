package product

import (
	//"ecommerce/database"
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"net/http"
	"strconv"
)

type UpdateProduct struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	pID, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please send a valid product ID", http.StatusBadRequest)
		return
	}

	var updatedProduct UpdateProduct

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Please sent valid JSON data", http.StatusBadRequest)
		return
	}

	_, err = h.productRepo.Update(repo.Product{
		ID: pID,
		Name: updatedProduct.Name,
		Description: updatedProduct.Description,
		Price: updatedProduct.Price,
		ImgUrl: updatedProduct.ImgUrl,
	})

	if err != nil {
		util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, "Updated successfully", http.StatusCreated)
}
