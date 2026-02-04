package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(getProducts))
	mux.Handle("POST /create-products", http.HandlerFunc(createProduct))

	globalRouterHandler := globalRouter(mux)
	fmt.Println("server running on port 8080")
	err := http.ListenAndServe(":8080", globalRouterHandler)
	if err != nil {
		fmt.Println("Error in starting the server", err)
	}
}

func init() {
	prd1 := Product{
		ID:          1,
		Name:        "Mango",
		Description: "Mango is tasty",
		Price:       63.4,
		ImgUrl:      "https://example.com/images/mango.jpg",
	}

	prd2 := Product{
		ID:          2,
		Name:        "Apple",
		Description: "Fresh and crispy apple",
		Price:       50.0,
		ImgUrl:      "https://example.com/images/apple.jpg",
	}

	prd3 := Product{
		ID:          3,
		Name:        "Banana",
		Description: "Sweet yellow banana",
		Price:       20.5,
		ImgUrl:      "https://example.com/images/banana.jpg",
	}

	prd4 := Product{
		ID:          4,
		Name:        "Orange",
		Description: "Juicy orange",
		Price:       35.0,
		ImgUrl:      "https://example.com/images/orange.jpg",
	}

	prd5 := Product{
		ID:          5,
		Name:        "Pineapple",
		Description: "Tropical pineapple",
		Price:       80.0,
		ImgUrl:      "https://example.com/images/pineapple.jpg",
	}

	prd6 := Product{
		ID:          6,
		Name:        "Strawberry",
		Description: "Fresh strawberries",
		Price:       120.0,
		ImgUrl:      "https://example.com/images/strawberry.jpg",
	}

	ProductList = append(ProductList, prd1, prd2, prd3, prd4, prd5, prd6)
}
