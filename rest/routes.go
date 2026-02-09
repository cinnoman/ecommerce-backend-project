package rest

import (
	"ecommerce/rest/handlers"
	"ecommerce/rest/middleware"
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
			handlers.GetProduct,
		),
	))

	mux.Handle("PUT /products/{id}", mngr.With(
		http.HandlerFunc(
			handlers.UpdateProduct,
		),
	))

	mux.Handle("DELETE /products/{id}", mngr.With(
		http.HandlerFunc(
			handlers.DeleteProduct,
		),
	))
}
