package routers

import (
	"encoding/json"
	"net/http"

	"github.com/lucassingh/go-backend/db"
	"github.com/lucassingh/go-backend/models"
)

/*UpdateProfile user*/
func UpdateProfile(w http.ResponseWriter, r *http.Request) {

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Incorrect data from user"+err.Error(), 400)
		return
	}

	var status bool

	status, err = db.UpdateRegister(t, IDUser)

	if err != nil {
		http.Error(w, "An error ocurred when try update user profile. Try again later..."+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "We coudnt update user register"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
