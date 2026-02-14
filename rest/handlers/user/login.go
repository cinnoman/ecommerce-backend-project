package user

import (
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

type credential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var reqLogin credential

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)
	if err != nil {
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
	}

	user, err := h.userRepo.Find(reqLogin.Email, reqLogin.Password)
	if err != nil {
		util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if user == nil {
		util.SendError(w, "Invalid Credentials", http.StatusBadRequest)
		return
	}

	accessToken, err := util.CreateJwt(h.cnf.JwtSecretKey, util.Payload{
		Sub:         user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		IsShopOwner: user.IsShopOwner,
	})

	if err != nil {
		util.SendError(w, "Not authorized", http.StatusMethodNotAllowed)
		return
	}

	util.SendData(w, accessToken, http.StatusOK)
}
