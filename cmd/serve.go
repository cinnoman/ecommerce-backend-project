package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	mngr := middleware.NewManager()
	mngr.Use(middleware.Logger, middleware.Hudai)


	mux.Handle("GET /noman", mngr.With(
		http.HandlerFunc(
			handlers.Test,
		),
	))

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

	globalRouterHandler := global_router.GlobalRouter(mux)
	fmt.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", globalRouterHandler)
	if err != nil {
		fmt.Println("Error in starting the server", err)
	}
}
