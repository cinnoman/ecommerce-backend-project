package database

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

var ProductList []Product

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