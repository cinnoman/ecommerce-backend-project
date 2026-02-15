package repo

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          int    `db:"id" json:"id"`
	FirstName   string `db:"first_name" json:"firstName"`
	LastName    string `db:"last_name" json:"lastName"`
	Email       string `db:"email" json:"email"`
	Password    string `db:"password" json:"password"`
	IsShopOwner bool   `db:"is_shop_owner" json:"isShopOwner"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	Find(email, pass string) (*User, error)
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(user User) (*User, error) {
	query := `
	INSERT INTO users (
		first_name,
		last_name,
		email,
		password,
		is_shop_owner
	) VALUES (
		:first_name,
		:last_name,
		:email,
		:password,
		:is_shop_owner
	)
	RETURNING id
	`

	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&user.ID)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("failed to insert user")
	}

	return &user, nil
}

func (r *userRepo) Find(email, pass string) (*User, error) {
	var user User
	query := `
	SELECT id, first_name, last_name, email, password, is_shop_owner
	FROM users
	WHERE email = $1 AND password = $2
	LIMIT 1
	`

	err := r.db.Get(&user, query, email, pass)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
