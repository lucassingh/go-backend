package routers

import (
	"encoding/json"
	"net/http"

	"github.com/lucassingh/go-backend/db"
	"github.com/lucassingh/go-backend/models"
)

/*Register put a register user in DB*/
func Register(w http.ResponseWriter, r *http.Request) {

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error data recived"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Password must be 6 characters at least", 400)
		return
	}

	_, find, _ := db.CheckUserExist(t.Email)

	if find == true {
		http.Error(w, "A user with this email already exists", 400)
		return
	}

	_, status, err := db.InsertRegister(t)

	if err != nil {
		http.Error(w, "An error occurred while attempting user registration"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Failed to insert user record"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
