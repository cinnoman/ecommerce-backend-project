package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser database.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		http.Error(w, "Please sent valid JSON data", http.StatusBadRequest)
		return
	}

	usr := newUser.Store()
	util.SendData(w, usr, http.StatusCreated)
}
