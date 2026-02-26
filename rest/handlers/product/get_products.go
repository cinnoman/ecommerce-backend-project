package product

import (
	"ecommerce/domain"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	reqQuery := r.URL.Query()

	pageAsStr := reqQuery.Get("page")
	limitAsStr := reqQuery.Get("limit")

	page, _ := strconv.Atoi(pageAsStr)
	limit, _ := strconv.Atoi(limitAsStr)

	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	prdList := make(chan []*domain.Product)
	go func() {
		productList, err := h.svc.List(page, limit)
		if err != nil {
			util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		prdList <- productList
	}()

	count := make(chan int)
	go func() {
		cnt, err := h.svc.Count()
		if err != nil {
			util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		count <- cnt
	}()

	
	productList := <-prdList
	cnt := <- count

	util.SendPage(w, productList, page, limit, cnt)
}
