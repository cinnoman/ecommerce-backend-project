package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)


type credential struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var reqLogin credential

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)
	if err != nil {
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
	}

	user := database.Find(reqLogin.Email, reqLogin.Password)
	if user == nil {
		util.SendError(w, "Invalid Credentials", http.StatusBadRequest)
		return
	}

	util.SendData(w, user, http.StatusOK)
}