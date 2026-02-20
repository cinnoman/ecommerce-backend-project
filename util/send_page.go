package util

import "net/http"

type PaginatedData struct {
	Data       any        `json:data`
	Pagination Pagination `json:pagination`
}

type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"totalItems"`
	TotalPages int `json:"totalPages"`
}

func SendPage(w http.ResponseWriter, data any, page, limit, cnt int) {

	paginatedData := PaginatedData{
		Data: data,
		Pagination: Pagination{
			Page:       page,
			Limit:      limit,
			TotalItems: cnt,
			TotalPages: cnt / limit,
		},
	}

	SendData(w, paginatedData, http.StatusOK)
}
