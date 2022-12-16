package routers

import (
	"encoding/json"
	"net/http"

	"github.com/lucassingh/go-backend/db"
	"github.com/lucassingh/go-backend/jwt"
	"github.com/lucassingh/go-backend/models"
)

/*Login*/
func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("contente-type", "aplication/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "User or password are incorrect"+err.Error(), 400)
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
	}

	document, exist := db.TryLogin(t.Email, t.Password)

	if exist == false {
		http.Error(w, "User or password are invalid", 400)

		return
	}

	jwtKey, err := jwt.GenerateJWT(document)

	if err != nil {
		http.Error(w, "Cannot generate token"+err.Error(), 400)
	}

	response := models.ResponseLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "aplication/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}
