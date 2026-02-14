package repo

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"isShopOwner"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	Find(email, pass string) (*User, error)
}

type userRepo struct {
	userList []User
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (r *userRepo) Create(user User) (*User, error) {
	if user.ID != 0 {
		return &user, nil
	}

	user.ID = len(r.userList) + 1
	r.userList = append(r.userList, user)
	return &user, nil
}

// func (r *userRepo) Find(email, pass string) (*User, error) {
// 	for _, user := range r.userList {
// 		if email == user.Email && pass == user.Password {
// 			return &user, nil
// 		}
// 	}

// 	return nil, nil
// }

func (r *userRepo) Find(email, pass string) (*User, error) {
	for i := range r.userList {
		if email == r.userList[i].Email && pass == r.userList[i].Password {
			return &r.userList[i], nil
		}
	}
	return nil, nil
}

