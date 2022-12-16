package routers

import (
	"encoding/json"
	"net/http"

	"github.com/lucassingh/go-backend/db"
)

/*ViewProfile*/
func ViewProfile(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "must send id parameter", http.StatusBadRequest)
		return
	}

	profile, err := db.FindProfile(ID)

	if err != nil {
		http.Error(w, "Error ocurred to find user"+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "aplication/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(profile)
}
