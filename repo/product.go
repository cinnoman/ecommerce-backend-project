package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int     `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImgUrl      string  `json:"imgUrl" db:"img_url"`
}

type ProductRepo interface {
	Get(pId int) (*Product, error)
	List() ([]*Product, error)
	Create(newProduct Product) (*Product, error)
	Update(product Product) (*Product, error)
	Delete(pId int) error
}

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(p Product) (*Product, error) {
	query := `
			INSERT INTO products (
			name,
			description,
			price,
			img_url
		) VALUES (
				$1,
				$2,
				$3,
				$4
			)
			RETURNING id
	`

	row := r.db.QueryRow(query, p.Name, p.Description, p.Price, p.ImgUrl)
	err := row.Scan(&p.ID)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *productRepo) Get(id int) (*Product, error) {
	var prd Product

	query := `
		SELECT id, name, description, price, img_url
		FROM products
		WHERE id=$1;
	`

	err := r.db.Get(&prd, query, id)
	if err != nil {
		return nil, err
	}

	return &prd, nil
}


func (r *productRepo) List() ([]*Product, error) {

	prdList := []*Product{}

	query := `
		SELECT id, name, description, price, img_url
		FROM products;
	`

	err := r.db.Select(&prdList, query)
	if err != nil {
		return nil, err
	}

	return prdList, nil
}


func (r *productRepo) Update(p Product) (*Product, error) {
	query := `
		UPDATE products SET
		name=$1, description=$2, price=$3, img_url=$4
		WHERE id=$5
	`

	result, err := r.db.Exec(query, p.Name, p.Description, p.Price, p.ImgUrl, p.ID)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, sql.ErrNoRows
	}

	return &p, nil
}

func (r *productRepo) Delete(id int) error {
	query := `
		DELETE FROM products WHERE id=$1
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
