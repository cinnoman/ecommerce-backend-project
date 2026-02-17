package product

import (
	"ecommerce/domain"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

type ReqCreateProduct struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req ReqCreateProduct

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Please sent valid JSON data", http.StatusBadRequest)
		return
	}

	createdProduct, err := h.svc.Create(domain.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		ImgUrl:      req.ImgUrl,
	})

	if err != nil {
		util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, createdProduct, http.StatusCreated)
}
