package repo

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

type ProductRepo interface {
	Get(pId int) (*Product, error)
	List() ([]*Product, error)
	Create(newProduct Product) (*Product, error)
	Update(product Product) (*Product, error)
	Delete(pId int) error
}

type productRepo struct {
	productList []*Product
}

func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	getInitialProduct(repo)
	return repo
}

func (r *productRepo) Create(newProduct Product) (*Product, error) {
	newProduct.ID = len(r.productList) + 1
	r.productList = append(r.productList, &newProduct)
	return &newProduct, nil
}

func (r *productRepo) Get(pId int) (*Product, error) {
	for _, product := range r.productList {
		if product.ID == pId {
			return product, nil
		}
	}
	return nil, nil
}

func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}

func (r *productRepo) Update(product Product) (*Product, error) {
	for idx, p := range r.productList {
		if p.ID == product.ID {
			r.productList[idx] = &product 
		}
	}

	return &product, nil
}

func (r *productRepo) Delete(pId int) error {
	var tempList []*Product

	for _, product := range r.productList {
		if product.ID != pId {
			tempList = append(tempList, product)
		}
	}

	r.productList = tempList
	return nil
}

func getInitialProduct(r *productRepo) {
	prd1 := &Product{
		ID:          1,
		Name:        "Mango",
		Description: "Mango is tasty",
		Price:       63.4,
		ImgUrl:      "https://example.com/images/mango.jpg",
	}

	prd2 := &Product{
		ID:          2,
		Name:        "Apple",
		Description: "Fresh and crispy apple",
		Price:       50.0,
		ImgUrl:      "https://example.com/images/apple.jpg",
	}

	prd3 := &Product{
		ID:          3,
		Name:        "Banana",
		Description: "Sweet yellow banana",
		Price:       20.5,
		ImgUrl:      "https://example.com/images/banana.jpg",
	}

	prd4 := &Product{
		ID:          4,
		Name:        "Orange",
		Description: "Juicy orange",
		Price:       35.0,
		ImgUrl:      "https://example.com/images/orange.jpg",
	}

	prd5 := &Product{
		ID:          5,
		Name:        "Pineapple",
		Description: "Tropical pineapple",
		Price:       80.0,
		ImgUrl:      "https://example.com/images/pineapple.jpg",
	}

	prd6 := &Product{
		ID:          6,
		Name:        "Strawberry",
		Description: "Fresh strawberries",
		Price:       120.0,
		ImgUrl:      "https://example.com/images/strawberry.jpg",
	}

	r.productList = append(r.productList, prd1, prd2, prd3, prd4, prd5, prd6)
}
