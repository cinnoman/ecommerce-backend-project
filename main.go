package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

func handlerCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Header", "Content-Type")
	w.Header().Set("Access-Control-Allow-Method", "POST")
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

var ProductList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	handlerCors(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
	sendData(w, ProductList, 200)
}


func createProduct(w http.ResponseWriter, r *http.Request) {

	handlerCors(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
	var newProduct Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Please sent valid JSON data", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(ProductList) + 1

	ProductList = append(ProductList, newProduct)

	sendData(w, newProduct, 201)
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(getProducts))
	mux.Handle("OPTIONS /products", http.HandlerFunc(getProducts))
	mux.Handle("OPTIONS /create-products", http.HandlerFunc(createProduct))
	mux.Handle("POST /create-products", http.HandlerFunc(createProduct))

	fmt.Println("server running on port 8080")
	err := http.ListenAndServe(":8080", mux)
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
