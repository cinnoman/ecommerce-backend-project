package database

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"isShopOwner"`
}

var userList []User

func (u User) Store() User {
	if u.ID != 0 {
		return u
	}

	u.ID = len(userList) + 1
	userList = append(userList, u)
	return u
}

func Find(email, pass string) *User {
	for _, user := range userList {
		if email == user.Email && pass == user.Password {
			return &user
		}
	}

	return nil
}
