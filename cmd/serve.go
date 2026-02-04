package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))
	mux.Handle("GET /products/{id}", http.HandlerFunc(handlers.GetProductByID))

	globalRouterHandler := global_router.GlobalRouter(mux)
	fmt.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", globalRouterHandler)
	if err != nil {
		fmt.Println("Error in starting the server", err)
	}
}
