package repo

import (
	"database/sql"

	"ecommerce/domain"

	"github.com/jmoiron/sqlx"
	"ecommerce/product"
)

type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {
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

func (r *productRepo) Get(id int) (*domain.Product, error) {
	var prd domain.Product

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

func (r *productRepo) List(page, limit int) ([]*domain.Product, error) {

	offset := ((page -1) * limit) + 1

	prdList := []*domain.Product{}

	query := `
		SELECT id, name, description, price, img_url
		FROM products
		LIMIT $1
		OFFSET $2;
	`

	err := r.db.Select(&prdList, query, limit, offset)
	if err != nil {
		return nil, err
	}

	return prdList, nil
}


func (r *productRepo) Count() (int, error) {

	query := `
		SELECT COUNT(*) FROM products;
	`
	var count int

	err := r.db.Get(&count, query)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *productRepo) Update(p domain.Product) (*domain.Product, error) {
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
