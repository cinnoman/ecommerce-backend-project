package user

import (
	"ecommerce/domain"
	usrHandler "ecommerce/rest/handlers/user"
)

type Service interface {
	usrHandler.Service
}


type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Find(email, pass string) (*domain.User, error)
}






