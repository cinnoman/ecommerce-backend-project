package cmd

import (
	"ecommerce/handlers"
	"ecommerce/middleware"
	"net/http"
)

func InitRoutes(mux *http.ServeMux, mngr *middleware.Manager) {
	mux.Handle("GET /products", mngr.With(
		http.HandlerFunc(
			handlers.GetProducts,
		),
	))

	mux.Handle("POST /products", mngr.With(
		http.HandlerFunc(
			handlers.CreateProduct,
		),
	))

	mux.Handle("GET /products/{id}", mngr.With(
		http.HandlerFunc(
			handlers.GetProductByID,
		),
	))
}
